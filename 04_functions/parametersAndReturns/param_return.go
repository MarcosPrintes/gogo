package main

import "fmt"

/*
	=> Funções podem pegar varios argumentos e retornar valores
	=> Funções em Go podem retornar mais de um valor
	=> Funções que retornam mais de um valor são úteis por exemplo, quando quero
	retornar um valor ou um erro
	=> Toda função que retorna ao menos um valor deve terminar com return ou panic
	=> uma função pode ser definida com nome e nenhuma nome para os parametros, só tipos
	func f(int, string)
	=> Os valores retornados podem ser nomeados ou sem nome (named, unamed)
	=> Se uma função retorna 4 ou mais valores, é melhor retornar um slice se os tipos
	retornados forem do mesmo tipo ou usar um PONTEIRO PARA UMA STRUCT	se eles forem
	de diferentes tipos
*/

func main() {

	fmt.Println("result from Multiply3Numbers=> ", Multiply3Numbers(1, 2, 3))

	fmt.Println("return pointer struct =>", f1())
}

func Multiply3Numbers(a, b, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

type ReturnType struct {
	err      string
	succes   int
	complete bool
}

func f1() *ReturnType {
	response := ReturnType{
		err:      "error message",
		succes:   1,
		complete: true,
	}
	return &response
}
