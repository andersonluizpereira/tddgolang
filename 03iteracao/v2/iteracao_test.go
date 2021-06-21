package v2

import "testing"

func TestIteracao(t *testing.T) {
	repeticoes := Repetir("a", 10)
	esperado := "aaaaaaaaaa"
	if repeticoes != esperado {
		t.Errorf("esperado %s está diferente de repeticoes %s", esperado, repeticoes)
	}
}

func TestIteracaoNumeroZero(t *testing.T) {
	repeticoes := Repetir("a", 0)
	esperado := ""
	if repeticoes != esperado {
		t.Errorf("esperado %s está diferente de repeticoes %s", esperado, repeticoes)
	}
}
