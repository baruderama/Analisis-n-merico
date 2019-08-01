package main

import (
	"fmt"
	"math"
)

func main() {

	a, b := formulaCuadratica()
	c, d := formulaCuadratica2()
	fmt.Printf("hey primera positiva,segunda negativa\n")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("hey formula2 primera positiva,segunda negativa\n")
	fmt.Println(c)
	fmt.Println(d)

}

/*
func f(x) {
	xcuadrado := math.Pow(x, 2)
	return (5 * xcuadrado) - (4 * x) + 3

}
*/

func formulaCuadratica() (float64, float64) {
	a := 1.0
	b := math.Pow(9, 12)
	c := -3.0
	//dentroDeRaiz := (b * b) - (4 * a * c)
	//raiz := math.Sqrt(dentroDeRaiz)
	cuadraticaPositiva := ((-b) + (math.Sqrt((b * b) - (4 * a * c)))) / (2 * a)
	cuadraticaNegativa := ((-b) - (math.Sqrt((b * b) - (4 * a * c)))) / (2 * a)
	//fmt.Println(dentroDeRaiz)
	return cuadraticaPositiva, cuadraticaNegativa

}

func formulaCuadratica2() (float64, float64) {
	a := 1.0
	b := math.Pow(9, 12)
	c := -3.0
	pos := (-2 * c) / (b + (math.Sqrt((b * b) - (4 * a))))
	neg := (-2 * c) / (b - (math.Sqrt((b * b) - (4 * a))))
	return pos, neg
}
