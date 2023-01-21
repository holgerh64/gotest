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
	var erg int

	vt.AnzR += 1
	erg = 0

	return erg
}
