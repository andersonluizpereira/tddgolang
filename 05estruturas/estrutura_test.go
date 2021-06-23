package estruturas

import "testing"

func TestPerimetro(t *testing.T) {
	t.Run("Teste de perimetro", func(t *testing.T) {
		resultado := Perimetro(10.0, 10.0)
		esperado := 40.0
		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f ", resultado, esperado)
		}
	})
	
	t.Run("Teste de Área", func(t *testing.T) {
		resultado := Area(12.0, 6.0)
		esperado := 72.0
		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	})
}
