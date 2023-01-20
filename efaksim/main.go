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
	"math"
	"math/rand"
)

//-----------------------------------------------
// Der Typ Vokabel
//-----------------------------------------------

// Deklaration der Methoden-Typen
//type VokNrGetter func() int

// Die Struktur Vokabel
type Vokabel struct {
	Voknr  int     // Vokabeltext durch laufende Nummer abgebildet
	InFach int     // Fach in dem sich die Vokabel befindet
	E      float64 // Der sagenhafte E-Faktor
	AnzR   int     // Anzahl der Abfragen mit korrerktem Ergebnis
	AnzF   int     // Anzahl der Abfragen mit falschem Ergebnis

	// Deklaration der setter Methoden

	// Deklaration der getter Methoden

}

// Struktur, die Statistiken zur Vokabelliste enthält
type VokStat struct {
	vokanz int     // Anzahl der Vokabeln in der Liste
	minE   float64 // Minimaler E-Wert
	meanE  float64 // Mittlerer E-Wert
	maxE   float64 // Maximaler E-Wert
}

// Implementierung der Setter Methoden

// Implementierung der getter Methoden

// Der Einsprungpunkt
func main() {
	var numVok int
	var x int
	numVok = 100 // Anzahl der Vokabeln in der Liste

	var v1 []Vokabel // Slice für die Vokabeln vorbereiten
	var vs VokStat

	// Zufallszahlengenerator initialisieren.
	// Konstante erzeugt stets gleiche Reihen von Zufalsszahlen
	rand.Seed(1) // Konstante durch Systemzeit ersetzen für stets "neue" Zahlen

	// Speicherplatz für die Vokabelliste beschaffen
	v1 = make([]Vokabel, numVok)

	for x = 0; x < len(v1); x++ {
		iniVok(&v1[x], 2.5, x+1)
	}

	for x = 0; x < len(v1); x++ {
		prtVok(&v1[x])
	}

	vs = calcVokStat(v1)
	fmt.Println("E Stats:\nEmin, Emean, Emax")
	fmt.Println(vs.minE, vs.meanE, vs.maxE)
	fmt.Println("Vokanz=", vs.vokanz)
}

// Eine Struktur vom Type Vokabel initialisieren
func iniVok(vt *Vokabel, et float64, vnr int) {
	vt.AnzF = 0
	vt.AnzR = 0
	vt.E = zZahl(1.0, et)
	vt.Voknr = vnr
	vt.InFach = 5

}

// Berechnet elementare Statistiken der Vokabelliste
func calcVokStat(vt []Vokabel) VokStat {
	var tStat VokStat
	var x int
	var tEmin, tEmean, tEmax float64

	tEmin = math.MaxFloat64
	tEmax = 0.0
	tEmean = 0.0

	for x = 0; x < len(vt); x++ {
		if vt[x].E < tEmin {
			tEmin = vt[x].E
		}
		if vt[x].E > tEmax {
			tEmax = vt[x].E
		}

		tEmean += vt[x].E
	}
	tEmean = tEmean / float64(len(vt))
	tStat.minE = tEmin
	tStat.maxE = tEmax
	tStat.meanE = tEmean
	tStat.vokanz = len(vt)
	return tStat
}

// Testfunktion zur Ausgabe von Werten aus Strukt Vokabel
func prtVok(vt *Vokabel) {
	fmt.Println("Vokabelnummer: ", vt.Voknr)
	fmt.Println(vt.E)
}

// Gibt eine Zufallsazhl im Bereich min bis max aus
func zZahl(min float64, max float64) float64 {
	var tmp float64

	tmp = 0.0

	tmp = rand.Float64()*(max-min) + min

	return tmp
}
