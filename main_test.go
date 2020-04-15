package main

import (
	"testing"

	"github.com/kevencript/challange-ci/soma"
)

func TestSoma(t *testing.T) {

	total := soma.SomarValor()

	if total != 10 {
		t.Error("Valor é diferente de 10!")
	}

}
