package testingtsupport_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
