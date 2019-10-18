package main

import "fmt"

/*
	=> Funções que possuem valores de retorno nomeadas podem inferir um valor para as mesmas
	durante a execução do bloco da função e simplismente dar o retorno vazia. (return)

	=> Utilizar variáveis de retorno nomeadas, geram código mais limpo e autdodocumentado


*/

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	a, s, x := mFunc1Named(1, 4)
	v, b, n := mFunc1Unamed(1, 4)
	fmt.Printf("%d %d %d", a, s, x)
	fmt.Printf("%d %d %d", v, b, n)
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

func mFunc1Named(a, b int) (sum, product, diff int) {
	sum = a + b
	product = a * b
	diff = a - b

	return
}

func mFunc1Unamed(a, b int) (int, int, int) {
	sum := a + b
	product := a * b
	diff := a - b

	return sum, product, diff
}

func Mysqrt() (err error) {

	var n float64 = 34.3 * 2
	if n < 0 {
		return err
	}

	return nil
}
