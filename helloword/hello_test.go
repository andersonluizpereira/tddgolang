package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola()
	esperado := "ola mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}


