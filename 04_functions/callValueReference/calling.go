package main

import "fmt"

/*
Chamada por valor
 - O jeito padrão do Go é passar uma variável como parametro para uma função, é feita uma
 cópia da váriavel, a função altera o valor da cópia, a original permanece

 Chamada por referência
 - Para alterar o valor original de uma variável, é preciso que uma função receba um ponteiro
 e seja passado um endereço de memória (usando &)

=> channels, maps, interfaces, slices são passados por referencia, mas o ponteiro fica omitdio
por padrão
*/

func main() {

	f1(2)

	c := 3
	f2(&c) //passando um endereço de memória
}

func f1(a int) {
	fmt.Println("variable a:", a)
}

func f2(a *int) { //através do ponteiro, posso mudar o valor original da variável
	*a = 12
}
