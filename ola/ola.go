package main

import "fmt"

func Ola(nome string) string {
	if nome == ""{
		nome = "mundo"
	}

    return "Olá, " + nome
}

func main() {
    fmt.Println(Ola("Marcel"))
}
