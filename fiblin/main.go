// ------------------------------------------------------------------------
//
// # Nicht-rekursive Berechnung der Fibonacci-Folge
//
// ------------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 09.02.2023
//
// ------------------------------------------------------------------------
package main

import "fmt"

func main() {
	var f_2 float64
	var f_1 float64
	var f_t float64

	var x int

	f_2 = 1.0
	f_1 = 1.0
	f_t = 0.0
	x = 0

	for x = 0; x < 1000; x++ {
		if x == 0 || x == 1 {
			fmt.Println("fibonacci(", x, ")=", f_2)
		} else {
			f_t = f_2 + f_1
			fmt.Println("fibonacci(", x, ")=", f_t)
			f_2 = f_1
			f_1 = f_t
		}
	}
	return
}
