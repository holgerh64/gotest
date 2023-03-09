// ------------------------------------------------------------------------
//
// Erzeugen einer Liste mit n-Einträgen mit eindeutigen Zufallszahlen
// Für die Erzeugung von Vokabellisten
//
// ------------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 09.03.2023
//
// ------------------------------------------------------------------------
package main

import "fmt"

func main() {
	var anz int

	fmt.Print("Länge der Liste:")
	_, _ = fmt.Scanf("%d", &anz)
	fmt.Println("Anzahl=", anz)
	arr := make([]int, anz)
	for x := 0; x < anz; x++ {
		arr[x] = 0
	}
	for x := 0; x < anz; x++ {
		fmt.Println("arr[", x, "]=", arr[x])
	}
	return
}
