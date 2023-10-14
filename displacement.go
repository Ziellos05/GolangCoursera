package main

import "fmt"

func MainDisplacement() {

	var a float64

	var vo float64

	var so float64

	var t float64

	fmt.Print("Ingrese la aceleración")
	fmt.Scan(&a)

	fmt.Print("Ingrese la velocidad inicial")
	fmt.Scan(&vo)

	fmt.Print("Ingrese la posición inicial")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Ahora ingrese el tiempo para calcular la posición del objeto")
	fmt.Scan(&t)
	fmt.Print("La posición del objeto es: ", fn(t))

}

func GenDisplaceFn(a float64, vo float64, so float64) func(val float64) float64 {
	return func(x float64) float64 {
		return 0.5*a*x*x + vo*x + so
	}
}
