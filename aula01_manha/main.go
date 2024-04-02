package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n####### Exercicio 1 #######")
	// Exercicio 1 - imprimindo o nome na tela

	var nome string
	var idade int

	nome = "Maria"
	idade = 30

	fmt.Println("Nome: ", nome)
	fmt.Println("Idade: ", idade)

	fmt.Println("\n####### Exercicio 2 #######")
	// Exercicio 2 - Clima
	var temperatura int = 30
	var umidade = 83.0
	var pressaoAtm = 1019

	fmt.Println("Temperatura mínima:", temperatura, "graus")
	fmt.Println("Umidade:", umidade, "%")
	fmt.Println("Pressão atmosferica:", pressaoAtm, "hPa")

	fmt.Println("\n####### Exercicio 3 #######")
	// Exercicio 3 - Declaracao de variaveis

	var lnome string = "Jaque"
	var sobrenome string = "Maria"
	lsobrenome := 6
	quantidadeDeFilhos := 2

	// var init idade
	var idade_3 int = 10
	// var licenca_para_dirigir = true
	var licencaParaDirigir = true
	// var estatura da pessoa int
	var estaturaDaPessoa int

	fmt.Println("As variaveis corretas sao: lnome, sobrenome, lsobrenome e quantidadeDeFilhos.\nOs valores atribuidos foram, respectivamente:", lnome, sobrenome, lsobrenome, quantidadeDeFilhos)
	fmt.Println("As variaveis corrigidas sao: `idade`, `licenca_para_dirigir`, e `estatura da pessoa int.\nOs seguintes valores foram atribuidos foram, respectivamente,", idade_3, licencaParaDirigir, estaturaDaPessoa)

	fmt.Println("\n####### Exercicio 4 #######")
	/*
		Exercicio 4 - Tipos de dados
		var sobrenome string = "Silva"
		var idade int = "25"
		boolean := "false"
		var salario string = 4585.90
		var nome string = "Fellipe"
	*/
	var sobrenome4 = "Silva"
	fmt.Println("corrigido o nome da variavel `sobrenome`. o valor atribuido é:", sobrenome4)

	var idade4 int = 25
	fmt.Println("corrigido o tipo `int` e o nome da variavel `idade`. o valor atribuido é:", idade4)

	var boolean = "false"
	fmt.Println("corrigido o tipo da variavel `boolean`. o valor atribuido é a string:", boolean)

	var salario float64 = 4585.90
	fmt.Printf("corrigido o tipo da variavel `string`. o valor atribuido é: %.2f\n", salario)

	var nome5 string = "Fellipe"
	fmt.Println("corrigido o nome da variavel `nome`:", nome5)

}
