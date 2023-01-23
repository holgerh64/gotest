// -------------------------------------------------------------------
//
// Pakage zur Simulation des User-Verhaltens
// Der User erhält eine Vokabel, die er fehlerfrei-, mit Hilfe,
// oder nicht wiederholt.
// Das Ergebnis dient dazu die Abfrage der Vokabel upzudatenen und
// den neuen E-Faktor, sowie das neue Ablagefach zu bestimmen.
//
// -------------------------------------------------------------------
//
// Autor: Holger Husemann
// Datum: 21.01.2023
//
// -------------------------------------------------------------------
//
// Exportierte Strukturen:
//
// Exportierte Funktionen:
// func VErgUser(vt *efaktor.Vokabel) int
//
// -------------------------------------------------------------------
package UserReaction

import (
	"math"

	"github.com/holgerh64/gotest/efaksim/efaktor"
)

func VErgUser(vt *efaktor.Vokabel) int {
	// Der Rückgabewert:
	// 5 – perfect response
	// 4 – correct response after a hesitation
	// 3 – correct response recalled with serious difficulty
	// 2 – incorrect response; where the correct one seemed easy to recall
	// 1 – incorrect response; the correct one remembered
	// 0 – complete blackout.
	var erg int                  // Ergebnis der Abfrage
	var anzR, anzF int           // Bisherige Anzahl richtiger und falscher Abfragen
	var oldE, ekorr float64      // alter E-Faktor, E-Korrekturfaktor
	var rawresp, anzkorr float64 // Zufallszahl, Korrekturfaktor für Anzahl der Abfragen
	// Variablen initiasisieren
	anzR = vt.AnzR
	anzF = vt.AnzF
	oldE = vt.E
	ekorr = 2.5-(oldE/1.2 - 1.083333333333333)
	rawresp = math.Log10(9.0*efaktor.ZZahl(0.0, 1.0) + 1.0)

	switch {
	case anzR <= anzF+24:
		anzkorr = 0.0
	case anzR > anzF+24:
		anzkorr = 1.0
	default:
		anzkorr = float64(anzR-anzF) / 24.0
	}

	geskorr := (ekorr + anzkorr+ rawresp)/3.0

	switch {
	case geskorr >= 0.9:
		erg = 5
	// break
	case geskorr >= 0.8 && geskorr <0.9:
		erg = 4
		// break
	case geskorr >= 0.6&& geskorr<0.8 :
		erg = 3
		// break
	case geskorr >= 0.4&& geskorr<0.6:
		erg = 2
		// break
	case geskorr > 0.3&&geskorr<0.4:
		erg = 1
		// break
	default:
		erg = 0
	}

	// Der Algoithmus
	// Lernrate
	// lerr := 0.1
	// Vergessensrat
	// forgr := 0.25
	// kleiner E-Fator => schwere Voabel
					// oft Falsch => schwere Vok

	return erg
}
