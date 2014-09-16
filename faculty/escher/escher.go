// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/fs"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	ns := faculty.Root.Refine("escher")
	ns.Grow("CircuitSourceDir", CircuitSourceDir{})
	ns.Grow("Lookup", Lookup{})
}

type Lookup struct{}

func (Lookup) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, dvalve string, dvalue interface{}) {
			if dvalve != "Address" {
				return
			}
			r := faculty.Root.Lookup(dvalue.(string))
			eye.Show("Circuit", r.(Circuit))
		}, 
		"Address", "Circuit",
	)
	return reflex
}

// CircuitSourceDir
type CircuitSourceDir struct{}

func (CircuitSourceDir) Materialize(matter *be.Matter) be.Reflex {
	return be.NewNounReflex(matter.Design.At(fs.Source{}).(Circuit).StringAt("Dir"))
}
