package main

import "fmt"

func ola(nome string) string{
	return "ola, "+nome
}

func main() {
	fmt.Println(ola("Anderson"))
}

