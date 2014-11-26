// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	// . "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func Init(srcdir string) {
	srcDir = srcdir
	faculty.Register("test.Match", be.NewMaterializer(&Match{}))
	faculty.Register("test.FilterAll", be.NewMaterializer(FilterAll{}))
	faculty.Register("test.Exec", be.NewMaterializer(Exec{}))
}

var srcDir string
