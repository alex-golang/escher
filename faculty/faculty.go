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

// Root is a global variable where packages can add gates as side-effect of being imported.
var Root = understand.NewFaculty()

// Func combines the name of a valve and an associated value.
type Func struct {
	Valve string
	Value interface{}
}

// Sentence is a collection of functionals, indexed by a rank that decreases with the 
// recency of the functional's update.
type Sentence tree.Tree // Rank:int -> Functional:Func

func (s Sentence) At(valve string) interface{} {
	for _, f := range s {
		if f.Valve == valve {
			return f.Value
		}
	}
	panic(7)
}

func (s Sentence) AtAsTree(valve string) tree.Tree {
	return tree.Make().Grow(valve, s.At(valve))
}

func (s Sentence) NumNonNil() (n int) {
	for _, f := range s {
		if f.Value == nil {
			n++
		}
	}
	return
}

// …
// The first entry is the most recent one.
type ShortCognize func(Sentence)

type ShortMemory struct {
	y []*think.Synapse
	recognizer ShortMemoryReCognizer
}

// 
func NewShortMemory(valve ...string) (think.Reflex, *ShortMemory) {
	reflex := make(think.Reflex)
	m := &ShortMemory{
		y: make([]*think.Synapse, len(valve)),
		recognizer: ShortMemoryReCognizer{
			recognize: make(map[string]*think.ReCognizer),
			memory: make(Sentence, len(valve)),
		},
	}
	for i, v := range valve {
		if _, ok := reflex[v]; ok {
			panic("duplicate valve")
		}
		reflex[v], m.y[i] = think.NewSynapse()
		m.recognizer.memory[i] = Func{Valve: v, Value: nil}
	}
	return reflex, m
}

func (m *ShortMemory) Attach(cognize ShortCognize) *ShortMemoryReCognizer {
	// Locking prevents individual completing Attach invocations 
	// from initiating cogntion, before all valves have been attached.
	m.recognizer.Lock()
	defer m.recognizer.Unlock()
	//
	m.recognizer.cognize = cognize
	for i, s := range m.recognizer.memory {
		s_ := s
		m.recognizer.recognize[s.Valve] = m.y[i].Attach(
			func(w interface{}) {
				m.recognizer.cognizeOn(s_.Valve, w)
			},
		)
	}
	return &m.recognizer
}

type ShortMemoryReCognizer struct {
	cognize ShortCognize
	recognize map[string]*think.ReCognizer
	sync.Mutex
	memory Sentence // most-recent to least recent
}

func (recognizer *ShortMemoryReCognizer) ReCognize(sentence Sentence) {
	ch := make(chan struct{})
	for _, s := range sentence {
		go func() {
			recognizer.recognize[s.Valve].ReCognize(s.Value)
			ch <- struct{}{}
		}()
	}
	for _ = range sentence {
		<-ch
	}
}

func (recognizer *ShortMemoryReCognizer) cognizeOn(valve string, value interface{}) {
	recognizer.Lock()
	i := recognizer.indexOf(valve)
	recognizer.memory[0], recognizer.memory[i] = recognizer.memory[i], recognizer.memory[0]
	recognizer.memory[0].Value = value
	r := make(Sentence, len(recognizer.memory))
	recognizer.Unlock()
	//
	copy(r, recognizer.memory)
	recognizer.cognize(r)
}

// indexOf returns the current index of valve in the most-recent-first order memory slice.
func (recognizer *ShortMemoryReCognizer) indexOf(valve string) int {
	for i, meme := range recognizer.memory {
		if meme.Valve == valve {
			return i
		}
	}
	panic(7)
}
