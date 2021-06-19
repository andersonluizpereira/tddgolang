package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola("Anderson")
	esperado := "ola, Anderson"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}


