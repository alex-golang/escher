// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package think provides a basis of four fundamental gates for manipulating thinkable images.
package think

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
)

func init() {
	ns := faculty.Root.Refine("think")
	//
	ns.Grow("Associate", Associate{})
	ns.Grow("Remember", Remember{})
	ns.Grow("Choose", Choose{})
	ns.Grow("Focus", Focus{})
	//
	ns.Grow("A", Associate{})
	ns.Grow("R", Remember{})
	ns.Grow("C", Choose{})
	ns.Grow("F", Focus{})
}

// Focus
type Focus struct{}

func (Focus) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer((&focus{}).Cognize, "From", "On", "When", "_")
	return reflex
}

type focus struct {
	sync.Mutex
	from Image
	on string
	when interface{}
}

func (x *focus) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "From":
		x.from = dvalue.(Image)
	case "On":
		x.on = dvalue.(string)
	case "When":
		x.when = dvalue
	case "_":
	default:
		panic("eh")
	}
	eye.Show(
		"_", 
		Image{
			"Focus": x.from.Copy().Cut(x.on),
			"When": x.when,
		},
	)
}
