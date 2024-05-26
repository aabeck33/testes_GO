package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com receptor (produto)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    12.32,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Caneta", 15.89, 0.05}
	fmt.Println(produto2, produto2.precoComDesconto())
}