// -----------------------------------------------------------------------------
//
// Simmulation des Vokabellernes mit 5 Föchern
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
// Lernquote
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
	"strconv"
)

//-----------------------------------------------
// Der Typ Vokabel
//-----------------------------------------------

// Deklaration der Methoden-Typen
//type VokNrGetter func() int

// Die Struktur Vokabel
type Vokabel struct {
	Voknr  int     // Vokabeltext durch laufebde Bummer abgebildet
	InFach int     // Fach in dem sich die Vokabel befindet
	E      float64 // Der sagenhafte E-Faktor
	AnzR   int     // Anzahl der Abfragen mit korrerktem Ergebnis
	AnzF   int     // Anzahl der Abfragen mit falschem Ergebnis

	// Deklaration der setter Methoden

	// Deklaration der getter Methoden

}

// Implementierung der Setter Methoden

// Implementierung der getter Methoden

// Der Einsprungpunkt
func main() {
	var v1 [100]Vokabel
	var x int

	fmt.Println("Hallo Welt")
	//prtVok(v1)

	for x = 0; x < 100; x++ {
		iniVok(&v1[x], 1.73, x+1)
	}
	for x = 0; x < 100; x++ {
		prtVok(&v1[x])
	}

}

func iniVok(vt *Vokabel, et float64, vnr int) {
	vt.AnzF = 0
	vt.AnzR = 0
	vt.E = et
	vt.Voknr = vnr
	vt.InFach = 5

}
func prtVok(vt *Vokabel) {
	fmt.Println("Vokabelnummer: " + strconv.Itoa(vt.Voknr))
	fmt.Println(&vt)
}
