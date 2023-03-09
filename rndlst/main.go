// ------------------------------------------------------------------------
//
// Erstelle eine Liste mit n eindeutigen ganzzahligen Zufallszahlen
// Als Test des Algorithmus für das Vokabelprogramm
//
// ------------------------------------------------------------------------
//
// Autor: Holgeer Husemann
// Datum: 09.03.2023
//
// ------------------------------------------------------------------------
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var anz int

	// Zufallszahlengenerator initialisieren
	// rand.Seed(time.Now().UnixNano())

	fmt.Print("Länge der Liste:")
	_, _ = fmt.Scanf("%d", &anz)
	lst := make([]int, anz)
	for x := 0; x < anz; x++ {
		lst[x] = -1
	}
	for x := 0; x < anz; x++ {
		lst[fndpos(lst, rand.Intn(anz-1))] = x
	}
	for x := 0; x < anz; x++ {
		fmt.Println("lst[", x, "]=", lst[x])
	}

	fmt.Println("Länge der Liste:", len(lst))
}

// Finde die nächste freie Position x mit lst[x] =-1
// Gebe die Position im Array zurück
// Wird keine freie Position gefunen, dann gebe -1 zurück
func fndpos(plst []int, sd int) int {
	var x, pos int
	var anz int
	x = 0
	anz = len(plst)
	pos = sd
	for x = 0; x < anz; x++ {
		if plst[pos] == -1 {

			return pos
		}
		if pos == anz-1 {
			pos = 0
		} else {
			pos++
		}
	}
	panic("No Position: Stop")
}
