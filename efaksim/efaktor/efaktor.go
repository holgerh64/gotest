// -------------------------------------------------------------------
//
// # Funktionen zur Berechnung des E-Faktor sind in diesem Paket
//
// -------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 21.01.2023
//
// -------------------------------------------------------------------
//
// Liste der Strukturen
package efaktor

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
	Vokanz int     // Anzahl der Vokabeln in der Liste
	MinE   float64 // Minimaler E-Wert
	MeanE  float64 // Mittlerer E-Wert
	MaxE   float64 // Maximaler E-Wert
}

// Implementierung der Setter Methoden

// Implementierung der getter Methoden
// Berechnet elementare Statistiken der Vokabelliste
func CalcVokStat(vt []Vokabel) VokStat {
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
	tStat.MinE = tEmin
	tStat.MaxE = tEmax
	tStat.MeanE = tEmean
	tStat.Vokanz = len(vt)
	return tStat
}

// Testfunktion zur Ausgabe von Werten aus Strukt Vokabel
func PrtVok(vt *Vokabel) {
	fmt.Println("Vokabelnummer: ", vt.Voknr)
	fmt.Println(vt.E)
}

// Gibt eine Zufallsazhl im Bereich min bis max aus
func ZZahl(min float64, max float64) float64 {
	var tmp float64

	tmp = 0.0

	tmp = rand.Float64()*(max-min) + min

	return tmp
}

// Eine Struktur vom Type Vokabel initialisieren
func IniVok(vt *Vokabel, et float64, vnr int) {
	vt.AnzF = 0
	vt.AnzR = 0
	vt.E = et
	vt.Voknr = vnr
	vt.InFach = 5

}

/*
Algorithm SM-2 used in the computer-based variant of the SuperMemo method
and involving the calculation of easiness factors for particular items:

1 - Split the knowledge into smallest possible items.

2 - With all items associate an E-Factor equal to 2.5.

3 - Repeat items using the following intervals:

	I(1):=1
	I(2):=6
	for n>2: I(n):=I(n-1)*EF
	where:
	I(n) – inter-repetition interval after the n-th repetition (in days),
	EF – E-Factor of a given item
	If interval is a fraction, round it up to the nearest integer.

4 - After each repetition assess the quality of repetition response in 0-5 grade scale:

	5 – perfect response
	4 – correct response after a hesitation
	3 – correct response recalled with serious difficulty
	2 – incorrect response; where the correct one seemed easy to recall
	1 – incorrect response; the correct one remembered
	0 – complete blackout.

5 - After each repetition modify the E-Factor of the recently repeated item according to the formula:

	EF’:=EF+(0.1-(5-q)*(0.08+(5-q)*0.02))
	where:
	EF’ – new value of the E-Factor,
	EF – old value of the E-Factor,
	q – quality of the response in the 0-5 grade scale.
	If EF is less than 1.3 then let EF be 1.3.

6 - If the quality response was lower than 3 then start repetitions for the item from the

	beginning without changing the E-Factor (i.e. use intervals I(1), I(2) etc. as if the
	item was memorized anew).

7 - After each repetition session of a given day repeat again all items that scored below

	four in the quality assessment. Continue the repetitions until all of these items score
	at least four.

The optimization procedure used in finding E-Factors proved to be very effective.
In SuperMemo programs you will always find an option for displaying the distribution of
E-Factors (later called the E-Distribution). The shape of the E-Distribution in a given
database was roughly established within few months since the outset of repetitions.
This means that E-Factors did not change significantly after that period and it is safe
to presume that E-Factors correspond roughly to the real factor by which the inter-repetition
intervals should increase in successive repetitions.
*/
func CalcEVal(vt *Vokabel, q int) float64 {
	var oldE, newE float64

	oldE = vt.E
	newE = oldE + (0.1 - float64(5-q)*(0.08+float64(5-q)*0.02))
	if newE < 1.3 {
		newE = 1.3
	}
	return newE
}
