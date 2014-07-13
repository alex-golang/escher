// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

// import (
// 	"fmt"
// 	"strconv"
// )

func SeeCircuit(src *Src) (cir *Circuit) {
	cir = &Circuit{}
	t := src.Copy()
	Space(t)
	cir.Name = Identifier(t) // empty-string identifier ok
	Space(t)
	if !t.TryMatch("{") {
		return nil
	}
	Space(t)
	for {
		switch q := SeePeerOrMatching(t).(type) {
		case *Matching:
			cir.Match = append(cir.Match, q)
			continue
		case *Peer:
			cir.Peer = append(cir.Peer, q)
			continue
		}
		break
	}
	Space(t)
	t.Match("}")
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return DesugarCircuit(cir)
}

func SeePeerOrMatching(src *Src) (v interface{}) {
	if p := SeePeer(src); p != nil {
		return p
	}
	if m := SeeMatching(src); m != nil {
		return m
	}
	return nil
}
