// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package circuit provides Escher gates for building dynamic cloud applications using the circuit runtime of http://gocircuit.org
package circuit

import (
	// "fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gocircuit/circuit/client"
	"github.com/gocircuit/escher/faculty"
)

// client *client.Client
func Init(discover string) {
	program = &Program{}
	if discover != "" {
		program.Client = client.DialDiscover(discover, nil)
	}

	rand.Seed(time.Now().UnixNano())

	ns := faculty.Root.Refine("circuit")
	ns.Grow("Process", Process{})
	ns.Grow("Docker", Docker{})
	// ns.Grow("Leaving", Leaving{})
	// ns.Grow("Joining", Joining{})
	// ns.Grow("Channel", Chan{})

	ns.Grow("ForkSpawn", ForkSpawn{})
	ns.Grow("ForkExit", ForkExit{})
	ns.Grow("ForkIO", ForkIO{})
}

// Program…
type Program struct {
	*client.Client
}

var program *Program

// ChooseID returns a unique textual ID.
func ChooseID() string {
	return strconv.FormatUint(uint64(rand.Int63()), 20)
}
