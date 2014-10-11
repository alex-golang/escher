// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit traversal and transformation.
package model

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("model.IO", be.NewGateMaterializer(IO{}))
	faculty.Register("model.Reservoir", be.NewGateMaterializer(&Reservoir{}))
	faculty.Register("model.DepthFirst", be.NewGateMaterializer(DepthFirst{}))
}
