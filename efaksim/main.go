// -----------------------------------------------------------------------------
//
// Simulation des Vokabellernes mit 5 Föchern
// Dabei wird der Lernerfolg jeder einzelner Vokabel mit dem E-Wert bewertet,
//
// -----------------------------------------------------------------------------
//
// E-Werte
// Die E-Werte sind in einem Array abgelegt. Der Lernende hat einen
// Fortschrittsfaktor, der eine Quote für neu zu lernende und zu vergessende
// Vokabeln festgelegt. Damit wird eine Abfrage durchgeführt und die E-Werte
// angepasst.
//
// Verteilung auf die Fächer und Verhalten der Vokabeln aufgrund ihrer
// Fächerzuordnung:
//
// Modellierung des Lernerfolgs:
// # Lernquote
//
// # Vergessensquote
//
// -----------------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 18.01.2023
//
// -----------------------------------------------------------------------------
package main

import (
	"fmt"
	"math/rand"

	"github.com/holgerh64/gotest/efaksim/UserReaction"
	"github.com/holgerh64/gotest/efaksim/efaktor"
)

// Der Einsprungpunkt
func main() {
	var numVok int
	var x int
	numVok = 100 // Anzahl der Vokabeln in der Liste

	var v1 []efaktor.Vokabel // Slice für die Vokabeln vorbereiten
	var vs efaktor.VokStat

	// Zufallszahlengenerator initialisieren.
	// Konstante erzeugt stets gleiche Reihen von Zufalsszahlen
	rand.Seed(1) // Konstante durch Systemzeit ersetzen für stets "neue" Zahlen

	// Speicherplatz für die Vokabelliste beschaffen
	v1 = make([]efaktor.Vokabel, numVok)

	for x = 0; x < len(v1); x++ {
		efaktor.IniVok(&v1[x], 2.5, x+1)
	}

	/* 	for x = 0; x < len(v1); x++ {
	   		efaktor.PrtVok(&v1[x])
	   	}
	*/
	for x = 0; x < 100; x++ {
		abfrage(v1)
		vs = efaktor.CalcVokStat(v1)
		// fmt.Println("E Stats:\nEmin, Emean, Emax")
		fmt.Println(vs.MinE, vs.MeanE, vs.MaxE)
		// vs.MinE,

		// ,vs.MaxE

		// fmt.Println("Vokanz=", vs.Vokanz)
	}
}

// -------------------------------------------------------------------
// Eine Abfrage durchführen (Ganze Vokabelliste,
// Auswahl im nächsten Schritt).
// Werte in den Strukturen werden neu berechet.
// -------------------------------------------------------------------
func abfrage(vt []efaktor.Vokabel) {
	var x int        // Zählvariablen
	var vl int       // Länge der Vokabelliste
	var resv int     // Resultat der User-Abfrage
	var newE float64 // Neuer E-Faktor

	vl = len(vt)

	for x = 0; x < vl; x++ {
		// User abfragen
		resv = UserReaction.VErgUser(&vt[x])
		// fmt.Println(resv)
		/* resv:  5 – perfect response
		4 – correct response after a hesitation
		3 – correct response recalled with serious difficulty
		2 – incorrect response; where the correct one seemed easy to recall
		1 – incorrect response; the correct one remembered
		0 – complete blackout. */
		// Neuen E-Faktor berechnen
		//func CalcEVal(vt *Vokabel, q int) float64
		newE = efaktor.CalcEVal(&vt[x], resv)
		// Werte updaten
		vt[x].E = newE
		if resv > 3 {
			vt[x].AnzR++
		} else {
			vt[x].AnzF++
		}
	}
	return
}
