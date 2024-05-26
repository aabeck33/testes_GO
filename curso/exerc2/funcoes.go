package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	fmt.Println("Esse é o início de tudo!")
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string){
	fmt.Printf("Parâmetro 1: %s e Parâmetro 2: %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4 (p1, p2 string) string {
	return fmt.Sprintf("F4: %s e %s.\n", p1, p2)
}

func f5 () (string, string){
	return "Retrno 1", "Retorno 2"
}

func f6(){
	debug.PrintStack()
}

func f7(){
	f6()
}

func f8(){
	f7()
}

func f9(p1,p2 int)(segundo int, primeiro int){
	segundo = p2
	primeiro = p1
	return
}

var soma1 = func(a, b int) int {
	return a + b
}

func multiplica(a, b int) int {
	return a * b
}

func exec(funcao func(int,int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func imprapr(aprovados ...string) {
	fmt.Println("Lista de aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func closure() func() {
	xx := 90
	var funcao = func() {
		fmt.Println(xx)
	}
	return funcao
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func fp1(n int) {
	n++
}

func fp2(n *int) {
	//* é usado para desrefenciar, ou seja, ter acesso ao valor
	*n++
}

func main(){
	defer fmt.Print("Esse é o fim de tudo!") //aguarda o fim de tudo para executar essa linha
	f1()
	f2("Parâmetro 1", "Parâmetro 2")
	r3, r4 := f3(), f4("Parâmetro 1", "Parâmetro 2")
	fmt.Println(r3+"\n", r4)
	r51, r52 := f5()
	r53, _ := f5()
	_, r54 := f5()
	fmt.Println("F5: ", r51, r52, r53, r54)
	f8()

	x,y := f9(2,5)
	fmt.Println(x, y)

	fmt.Println(soma1(2,9))
	soma2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(soma2(3,9))

	result1 := exec(multiplica, 39, 7)
	fmt.Println(result1)

	fmt.Println(media(1,2,3,4,5,6,10))
	fmt.Println(media(3,4))
	fmt.Println(media(3344.4,34233.3453,435.674,635.5))

	lista_aprovados := []string{"Alvaro", "Adriano", "Beck"}
	imprapr(lista_aprovados...)

	xx := 20
	fmt.Println(xx)
	imprimeX := closure()
	imprimeX()

	result2, _ := fatorial(5)
	fmt.Println(result2)
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	fp := 1
	fp1(fp)
	fmt.Println(fp)
	//Passando um ponteiro para a própria variável que será afetada
	fp2(&fp)
	fmt.Println(fp)
}