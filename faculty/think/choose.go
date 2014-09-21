// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	// "fmt"
	"math/rand"
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
)

// Choose
type Choose struct{}

func (Choose) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer((&choose{}).Cognize, "When", "From", DefaultValve)
	return reflex
}

type choose struct {
	sync.Mutex
	from Image // image from which a child is being chosen
	when interface{} // signal for choice
}

func (x *choose) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "From":
		x.from = dvalue.(Image)
	case "When":
		x.when = dvalue
	case DefaultValve:
	default:
		panic("eh")
	}
	j := rand.Intn(x.from.Len())
	for i, key := range x.from.Names() {
		if i != j {
			continue
		}
		eye.Show(
			DefaultValve, 
			Image{
				"When": x.when,
				"Choice": x.from[key],
			},
		)
	}
}
