package main

import (
	"fmt"
)

func Adiciona(a, b int) int {
	return a + b
}

func Subtrai(a, b int) int {
	return a - b
}


func main() {
	soma := Adiciona(5, 3)
	subtracao := Subtrai(10, 4)

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
}
