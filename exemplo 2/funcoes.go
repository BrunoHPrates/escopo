package main

func boasVindas(nome string) string {
	return	"Bem-vindo(a) " + nome + "! Você está na " + nomeEscola
}

func statusMaioridade(idade int) string {	
	if idade >= 18 {
		return "Você é maior de idade."
	} else {
		return "Você é menor de idade."
	}
}