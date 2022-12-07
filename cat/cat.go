// Liest von stdin und schreibt auf stdout. Zeilenweise
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithProg := os.Args

	// Array mit unbekannter anzahl von Elementen
	var text []string
	var x int = 0

	if len(os.Args) != 2 {
		fmt.Println("\nUSAGE: Cat <file to print>. Found: ", argsWithProg)
		os.Exit(0)
	}
	//fmt.Println("Open: ", os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())

		fmt.Println(text[x])
		x++
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Println("\n\nNumber of Textlines: ", len(text)+1)
}
