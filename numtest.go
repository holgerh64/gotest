//-------------------------------------------------------------------------
//
// Das infamous numtest-Programm. Es prüft die Genauigkeit atomarer Typen.
// Ursprünglich aus einem Modulo 2 Lehrwerk der ehemaligen DDR.
// Gekauft Ende 1980er in Berlin, Haoutstadt der DDR.
//
//-------------------------------------------------------------------------
//
// Autor: Holger HUsemann
// Datum: 07.08.2022
//
//-------------------------------------------------------------------------

// Deklaration der Package. Diese enthält die Haupteinsprungfunktion main.
package main

// Einbindung des Ein-Ausgabepacketes aud der Standard-Library.
import "fmt"

// Die Haputfunktion
func main() {
	// Drei Variablen werden deklariert, ein Zähler x als normaler Integer,
	// a, b als doppelt genaue Real-Typen
	var x int
	var a, b float64

	fmt.Println("Hallo Welt!")
	a = 1.0
	x = 0

	// Die universelle for-Schleife kann auch while- und do-Schleifen
	// abbilden. In diesem Fall eine while-Schleife.
	for a+1.0 != 1.0 {
		a /= 2.0
		x++
	}

	// Wir haben die Zahl gefunden, die zu 1.0 addiert nicht von 1.00
	// unterschieden werden kann.
	fmt.Println(a, "+1.0==1.0, nach:", x, " Durchläufen.")

	// Nun berechnen wir die größte Zahl, die durch 2 dividiert, vom
	// System als 0.0 angesehen wird.
	a = 1.0
	b = a
	x = 0
	for a != 0.0 {
		x++
		b = a
		a /= 2.0

	}

	// Die Zahl und die Anzahl der /2-Operationen werden ausgegeben.
	fmt.Println(b, "/2.0=0 nach:", x, " Durchläufen.")
}
