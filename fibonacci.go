// -------------------------------------------------------------------
//
// Rekursive Definition der Fibonacci-Zahlen.
//
// -------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 07.08.2022
//
// -------------------------------------------------------------------
package main

import (
	"fmt"
	"log"
	"time"
)

// Rekursive Bestimmung der n.ten Fibonacci-Zahl
func fibo(n int) int64 {
	if n == 0 || n == 1 {
		return (1)

	} else {
		return (fibo(n-1) + fibo(n-2))
	}
}

func main() {
	var number int

	// Einlesen der oberen Grenze
	fmt.Print("Bis zu welcher Fibonacci-Zahl soll berechnet werden?:")
	_, err := fmt.Scanf("%d", &number)

	// Bei fehlerhafter EIngabe, Programmende
	if err != nil {
		log.Fatal(err)
	}

	// Zeitname zu Beginn der Berechnung
	start := time.Now()

	for x := 1; x <= number; x++ {
		fmt.Println("Die ", x, ".te Fibonacci-Zahl ist:", fibo(x))
	}

	// Dauer der Berechnung feststellen
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Verstrichene Zeit:", elapsed)
}
