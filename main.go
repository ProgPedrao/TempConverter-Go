package main

import "fmt"

func main() {
	fmt.Print("Digite o valor de °Kelvin: ")

	var tempKelvin float64
	fmt.Scanln(&tempKelvin)

	tempC := tempKelvin - 273

	fmt.Printf("%g°K = %g°C", tempKelvin, tempC)
}
