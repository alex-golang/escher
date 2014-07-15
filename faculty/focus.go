// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"sync"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/understand"
)

// Time—>SentenceFunctional
type Sentence tree.Tree

// "Valve"—>string, "Value"—>interface{}
type SentenceFunctional tree.Tree

// Valve—>MemoryFunctional
type Memory tree.Tree

func (m Memory) At(valve string) MemoryFunctional {
	return tree.Tree(m).At(valve).(MemoryFunctional)
}

// "Valve"—>string, "Value"—>interface{}, "Age"—>int
type MemoryFunctional tree.Tree

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Sentence)

// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
type Eye struct {
	synapse map[string]*think.Synapse
	attention EyeReCognizer
}

type EyeReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	age int
	memory Memory
}

// NewEye creates a new short-term memory mechanism.
func NewEye(valve ...string) (think.Reflex, *Eye) {
	reflex := make(think.Reflex)
	m := &Eye{
		synapse: make(map[string]*think.Synapse),
		attention: EyeReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: make(Memory),
		},
	}
	for _, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("two valves, same name")
		}
		reflex[v], m.synapse[v] = think.NewSynapse()
		m.attention.memory.Grow(v, tree.Plant("Valve", v).Grow("Value", nil).Grow("Age", 0))
	}
	return reflex, m
}

func (m *Eye) Focus(cognize ShortCognize) *EyeReCognizer {
	// Locking prevents individual competing Focus invocations 
	// from initiating cogntion before all valves/synapses have been attached.
	m.attention.Lock()
	defer m.attention.Unlock()
	m.attention.cognize = cognize
	for v_, _ := range m.attention.memory {
		v := v_
		m.attention.recognize[v] = m.synapse[v].Focus(
			func(w interface{}) {
				m.attention.cognize(v, w)
			},
		)
	}
	return &m.attention
}
