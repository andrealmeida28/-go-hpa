package main

import (
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	esperado := "<b>Code.education Rocks!</b>"
	obtido := greeting("Code.education Rocks!")
	if !strings.Contains(obtido, esperado) {
		t.Errorf("Negrito nao encontrado. Obtido: %s, esperado: %s.", obtido, esperado)
	}
}
