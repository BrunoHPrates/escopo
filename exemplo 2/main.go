package main

import "fmt"

var nomeEscola = "Escola Técnica SENAI"

func main() {
	nome := "Bruno"
	idade := 16

	mensagem := boasVindas(nome)
	fmt.Println(mensagem)

	status := statusMaioridade(idade)
	fmt.Println(status)
}