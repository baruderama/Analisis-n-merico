package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := puntoFijo()
	fmt.Printf("este es el valor\n")
	fmt.Println(a)
	fmt.Printf("este es el error\n")
	fmt.Println(b * 100)

}

func puntoFijo() (float64, float64) {
	x := 0.0 //valor inicial
	err := 0.00000000000000000000001
	for {

		errorActual := math.Abs((g(x) - x) / g(x))
		//fmt.Printf("error\n")
		//fmt.Println(errorActual)

		if errorActual <= err {
			return g(x), errorActual //si el error es igual o menor al requerido (err) retorna el
			//valor obtenido y el error actual
		}
		x = g(x) //si el error no es el requerido x va a ser el valor que de la func g(x)
		//fmt.Printf("hey valor inicial\n")
		//fmt.Println(x)

	}
}

func g(x float64) float64 {
	aux := math.Pow(x, 2)
	//fmt.Println((aux - 2) / 5)//funación g(x) que me invente yo
	return (aux - 2) / 5
}
