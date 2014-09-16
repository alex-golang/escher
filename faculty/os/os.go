// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func Init(a string) {
	args = make(map[string]string) // n1=v1,n2=v2
	for _, p := range strings.Split(a, ",") {
		if p == "" {
			continue
		}
		nv := strings.Split(p, "=")
		if len(nv) != 2 {
			panic("command-line argument syntax")
		}
		v, err := url.QueryUnescape(nv[1])
		if err != nil {
			panic(err)
		}
		args[nv[0]] = v
		log.Printf("Argument %s=%s", nv[0], v)
	}
	ns := faculty.Root.Refine("os")
	ns.Grow("Arg", Arg{})
	ns.Grow("Env", Env{})
	ns.Grow("Exit", Exit{})
	ns.Grow("Fatal", Fatal{})
	ns.Grow("Stdin", Stdin{})
	ns.Grow("Stdout", Stdout{})
	ns.Grow("Stderr", Stderr{})
	//
	ns.Grow("LookPath", LookPath{})
	ns.Grow("Process", Process{})
	ns.Grow("ForkCommand", ForkCommand{})
	ns.Grow("ForkExit", ForkExit{})
	ns.Grow("ForkIO", ForkIO{})
	//
}

var args map[string]string

// Arg
type Arg struct{}

func (Arg) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			if valve != "Name" {
				return
			}
			n, ok := value.(string)
			if !ok {
				panic("non-string name perceived by os.arg")
			}
			eye.Show("Value", args[n])
		}, 
		"Name", "Value",
	)
	return reflex
}

// Env
type Env struct{}

func (Env) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			if valve != "Name" {
				return
			}
			n, ok := value.(string)
			if !ok {
				panic("non-string name perceived by os.env")
			}
			ev := os.Getenv(n)
			log.Printf("Environment %s=%s", n, ev)
			eye.Show("Value", ev)
		},
		"Name", "Value",
	)
	return reflex
}

// Exit
type Exit struct{}

func (Exit) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			os.Exit(value.(int))
		}, 
		"_",
	)
	return reflex
}

// Fatal
type Fatal struct{}

func (Fatal) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			log.Fatalln(value)
		}, 
		"_",
	)
	return reflex
}

// LookPath
type LookPath struct{}

func (LookPath) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer(
		func(eye *be.Eye, valve string, value interface{}) {
			if valve != "Name" {
				return
			}
			p, err := exec.LookPath(value.(string))
			if err != nil {
				log.Fatalf("no file path to %s", value.(string))
			}
			eye.Show("_", p)
		},
		"Name", "_",
	)
	return reflex
}
