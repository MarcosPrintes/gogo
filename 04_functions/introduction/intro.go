package main

import "fmt"

/*
	Go é compilado na ordem em que as funções são escritas, por questão de legibilidade
	é melhor começar pela main() e seguir pela ordem de chamada

	Tipos:

	-Funçõe normais com um identidicador
	-Funções anônimas, ou funções lambda
  -Methods

  => Go não suporta (ainda) sobrecarga de funções, que é quando tenho duas funções com mesmo nome
  mas com diferentes argumentos ou tipos de retorno

  => Para funções fora do Go, basta adicionar o nome e assinatura, sem o corpo

  => Podem ser usadas com tipos

  => São valores de primeira classe, logo podem ser atribuidas á variáveis a:= f1()

  => Uma função não pode ser declarada dentro de outra, mas isso pode ser imitado
  através de funções anônimas
  => Go não trabalha com generics, usar interfaces e um tipo switch, porém isso aumenta
  a complexidade do código e performance
*/

func main() {
	// uma função pode chamar/invocar outra
	printSomething()

	/*
		uma função pode ser passada como parametro para outra
		caso haja argumentos em f1, f2 deve retornar a mesma quantidade de argumentos
		que f1 precisa
	*/
	f1(f2())

	a := m1(2, 2)
	fmt.Println("something printed", a)

}

func m1(a, b int) int {
	fmt.Println("something printed", a+b)
	return 2
}

func m2(a int) int {
	return a
}

func printSomething() {
	fmt.Println("something printed")
}

func f1(a, b, c int) {
	fmt.Println("sum a + b + c: ", a+b+c)
}

func f2() (int, int, int) {
	return 1, 2, 3
}

//usando uma função como um tipo, o body é omitido
type mFunc func() int
