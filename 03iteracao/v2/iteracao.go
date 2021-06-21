package v2

func Repetir(caractere string, quantidadeRepeticoes int ) string {
	println(quantidadeRepeticoes)
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
