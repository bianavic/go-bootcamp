package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("\n####### Exercicio 1 - letras de uma palavra #######")

	word := "paralelepipedo"
	var countLetters int = len(word)
	var letters string

	fmt.Printf("\nThe word is %s. It contains %d letters\n", word, countLetters)

	for _, letter := range word {
		letters += string(letter) + "\n"
	}
	fmt.Println("The letters are:")
	fmt.Println(letters)

	fmt.Println("\n####### Exercicio 2 - emprestimo #######")

	// criterios
	fmt.Println("\nEnter the client's age:")
	var age int
	fmt.Scan(&age)

	fmt.Println("Is the client employed? (true/false):")

	var employed bool
	fmt.Scan(&employed)

	fmt.Println("How many years has the client been employed?")
	var yearsEmployed int
	fmt.Scan(&yearsEmployed)

	fmt.Println("What is the client's salary:")
	var salary float64
	fmt.Scan(&salary)

	// check (refactor to implement sth like)
	if age > 22 && employed && yearsEmployed > 1 {
		if salary > 100000 {
			fmt.Println("\nThe client can receive a loan without interest.")
		} else {
			fmt.Println("\nThe client can receive a loan with interest.")
		}
	} else {
		fmt.Println("\nThe client does not meet the criteria to receive a loan with interest.")
	}

	fmt.Println("\n####### Exercicio 3 - a que mês corresponde? #######")
	/*
		Abaixo coloquei 2 possíveis soluções para a atividade. Embora ambas cheguem ao mesmo resultado,
		optaria pela segunda solução, utilizando o pacote time.
		A solução 1 é mais simples e direta, e não requer a adição de pacotes adicionais, porém,
		o pacote facilita a formatação, lidando melhor com meses e números incorretos, trazendo mais consistencia e melhora da manutenção do codigo.
	*/

	/* solution 1
	var months = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Println(2, "de", months[2])
	*/

	// solution 2 - implementando o pacote Time
	monthNumber := 2
	month := time.Month(monthNumber)
	fmt.Printf("\n%d de %v\n", monthNumber, month)

	fmt.Println("\n####### Exercicio 4 - quantos anos tem... #######")

	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}
	fmt.Println("\nBejamin's age is", employees["Benjamin"])

	// older than 21
	count := 0
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}
	fmt.Printf("There are %d employees that have more than 21 years old\n", count)

	employees["Frederico"] = 25

	fmt.Print("Added to list: Frederico, 25 years old\n")

	delete(employees, "Pedro")
	fmt.Print("Deleted from list: Pedro\n")

	fmt.Println("\nUpdated employee list:")
	for name, age := range employees {
		fmt.Printf("%s: %d\n", name, age)
	}
}
