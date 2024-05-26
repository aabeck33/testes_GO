/*
Esse é o primeiro exercício de programação em GO!
*/

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func tipo(vr interface{}) string {
	switch vr.(type) {
	case int:
		return "Inteiro"
	case float32:
		return "Float 32"
	case float64:
		return "Float 64"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não mapeado"
	}
}
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2 //XOR - OU Exclusivo
	comprarSorv := trab1 || trab2

	return comprarSorv, comprarTV32, comprarTV50
}

func operar(pw float64, pq float64, op string) float64 {
	var result float64
	var pwi, pqi, resulti int
	switch op {
	case "soma":
		result = pw + pq
	case "subtrai":
		result = pw - pq
	case "multiplica":
		result = pw * pq
	case "divide":
		result = pw / pq
	case "modulo":
		pwi = int(pw)
		pqi = int(pq)
		resulti = pwi % pqi
		result = float64(resulti)
	case "E":
		pwi = int(pw)
		pqi = int(pq)
		resulti = pwi & pqi
		result = float64(resulti)
	case "OU":
		pwi = int(pw)
		pqi = int(pq)
		resulti = pwi | pqi
		result = float64(resulti)
	case "XOR":
		pwi = int(pw)
		pqi = int(pq)
		resulti = pwi ^ pqi
		result = float64(resulti)
	case "pot":
		result = m.Pow(pw, pq)
	case "max":
		result = m.Max(pw, pq)
	case "min":
		result = m.Min(pw, pq)
	}
	return (result)
}

func main() {
	const PI float64 = 3.141592653
	var raio = 3.2 // tipo float64 inferido pelo compilador
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	g, f, t := 5, false, true

	area := PI * m.Pow(raio, 2)

	fmt.Println("Esse é o primeiro exercício!")
	fmt.Print("A área da circunferência com raio ", raio, " é ", area, "\n")
	fmt.Printf("%d %d %d %x %v %t %t \n", a, b, c, d, g, f, t)
	fmt.Printf("O valor de PI é %.4f.", PI)
	//Tipos de variáveis
	// Inteiro sem sinal -> uint8 uint16 uint32 uint64
	fmt.Println("LIteral inteiro é ", reflect.TypeOf(320000))
	maximo := m.MaxInt64
	fmt.Println(maximo)
	var e byte = 255
	fmt.Println(e)
	// Interio com sinal -> int8 int16 int32 int64

	// Nùmeros reais -> float32 float64

	// boolean -> true false

	// tipo string
	var l string = "aa"
	// Tipo rune - referência ao código unicode
	var r rune = 'a' // código unicode - int32
	fmt.Println(r)
	// string de múltiplas linhas
	s2 := `Linha 1
	linha 2
	linha 3 ...`
	fmt.Println(s2)
	//Char é int32
	var k string = "a"

	fmt.Println(k, l)

	var m int
	var n float64
	var o float32
	var p string
	var q bool
	var z byte
	var s rune
	var h *int
	var j *string

	fmt.Printf("%v %v %v %q %v %v %v %v %v", m, n, o, p, q, z, s, h, j)

	//Conversões de tipo
	pe := 6.9
	te := 2
	fmt.Println(pe / float64(te))

	//Cuidado:
	fmt.Println("TEste " + string(97))

	//Int <-> String
	fmt.Println("Teste " + strconv.Itoa(97))
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	bo, _ := strconv.ParseBool("true") //true ou 1 - false ou qualquer outra coisa
	if bo {
		fmt.Println("Verdadeiro")
	}

	fmt.Println(operar(2, 3, "E"))

	var oo int
	oo = 20
	oo += g //oo = oo + g
	fmt.Println(oo, g)
	oo -= g
	g++
	fmt.Println(oo, g)
	oo /= g
	g--
	fmt.Println(oo, g)
	oo *= g
	fmt.Println(oo, g)
	oo %= 2
	oo, g = g, oo
	fmt.Println(oo, g)

	var d1, d2 time.Time
	d1 = time.Unix(0, 0)
	d2 = time.Unix(0, 0)
	d3 := time.UTC
	fmt.Println(d1, d2, d3, reflect.TypeOf(d3))

	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println(p1 == p2)

	sorvete, tv32, tv50 := compras(false, false)
	fmt.Printf("Sorvete: %t, TV32: %t, TV50: %t, Saudável: %t \n",
		sorvete, tv32, tv50, !sorvete)

	//ponteiros
	meuint := 1
	var point *int = nil //define um ponteiro int
	point = &meuint      //pega o endereço de mem da variável meuint

	*point++
	meuint++
	fmt.Println(point, *point, meuint, &meuint)
	//
	if pg := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10); pg > 5 {
		fmt.Println("Maior que 5")
		fmt.Println(pg)
	} else {
		fmt.Println("Menor que 5")
		fmt.Println(pg)
	}
	//fmt.Println(pg)

	ct := 1
	for ct <= 10 { //não exite while
		fmt.Println("Contador fora -> ", ct)
		ct++
	}

	for cf := 1; cf <= 10; cf++ {
		fmt.Println("Contador dentro ->", cf)
	}
	/*
		for {
			fmt.Println("Para sempre")
			time.Sleep(time.Second * 5)
		}
	*/
	fmt.Println(tipo("werw"))

	var notas [10]float64
	notas[0], notas[1], notas[2] = 7.5, 7.3, 3.5
	fmt.Println(notas)
	totaln := 0.0
	for cf := 0; cf < len(notas); cf++ {
		totaln += notas[cf]
	}
	media := totaln / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

	numeros := [...]int{1, 2, 3, 4, 56, 6, 7, 8, 9, 0, 90}
	for cf, numero := range numeros {
		fmt.Println(cf, numero)
	}
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	//Slice
	array1 := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}
	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	slice2 := array2[1:3]
	fmt.Println(slice2)
	slice3 := array2[:4]
	fmt.Println(slice3)
	slice4 := slice3[:2]
	fmt.Println(slice4)

	sl1 := make([]int, 20) //Slice com 20 elemenstos de um array interno de 20
	sl1[9] = 12
	fmt.Println(sl1)
	sl2 := make([]int, 20, 30) //Slice com 20 elementos de um array interno de 30
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl2 = append(sl2, 2, 5, 6, 7, 5, 3, 3, 578, 56, 45, 34)
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl3 := append(sl2, 2, 45, 34)
	sl4 := sl2[:30]
	sl5 := sl2[25]
	sl2[25] = 3454
	fmt.Println(sl2, sl3[25], sl4, reflect.TypeOf(sl4), sl5, reflect.TypeOf(sl5))

	//Map
	var aprovados map[int]string
	aprovados = make(map[int]string)
	aprovados[123] = "Alvaro"
	aprovados[345] = "Adriano"
	aprovados[678] = "Beck"
	fmt.Println(aprovados)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
		if cpf == 345 {
			delete(aprovados, cpf)
		}
	}
	fmt.Println(aprovados)

	funEsalario := map[string]float64{
		"Alvaro":  34324.34,
		"Adriano": 67755.43,
		"Beck":    23232.98,
	}
	funEsalario["Alvaro Beck"] = 344343.87
	for nome, salario := range funEsalario {
		fmt.Println(nome+": ", salario)
	}

	FuncPorLetra := map[string]map[string]float64{
		"A": {
			"Alvaro":  34324.34,
			"Adriano": 67755.43,
		},
		"B": {
			"Beck": 23232.98,
		},
	}
	fmt.Println(FuncPorLetra)
	for letra := range FuncPorLetra {
		for nome, salario := range FuncPorLetra[letra] {
			fmt.Println(letra, nome, salario)
		}
	}
}
