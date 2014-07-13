// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	// "fmt"
	"github.com/petar/maymounkov.io/escher/kit/record"
)

func SeeRecord(src *Src) (rec RecordDesign, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			rec, ok = nil, false
		}
	}()
	// println(fmt.Sprintf("R<=[%s]", src.String()))
	rec = RecordDesign(record.Make())
	t := src.Copy()
	t.Match("{")
	Space(t)
	for {
		q := t.Copy()
		Space(q)
		name, scope, ok := SeeField(q)
		if !ok {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		for _, w := range scope {
			(record.Record)(rec).Extend(name, w)
		}
		t.Become(q)
	}
	Space(t)
	t.TryMatch(",")
	Space(t)
	t.Match("}")
	src.Become(t)
	return rec, true
}
