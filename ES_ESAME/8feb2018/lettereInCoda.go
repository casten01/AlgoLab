//LETTERE IN CODA

package main

import (
	"fmt"
	"os"
	"strings"
)

func insertCoda(s string) (coda []rune) {
	coda = make([]rune, len(s))
	for _, v := range s {
		coda = append(coda, rune(v))
	}
	return coda
}

func spartisci(coda []rune) (S, C, V []rune) {
	for _, v := range coda {
		c := strings.ToUpper(string(v))
		if c == "A" || c == "E" || c == "I" || c == "O" || c == "U" {
			V = append(V, v)
		} else if c == "Y" || c == "J" || c == "W" || c == "K" || c == "X" {
			S = append(S, v)
		} else {
			C = append(C, v)
		}
	}
	return S, C, V
}

func extract(coda []rune) []rune {
	for i, v := range coda {
		i = 0
		fmt.Printf("%c", v)
		coda = coda[i:]
	}
	fmt.Println()
	return coda
}

func main() {
	var s string
	var codaS, codaC, codaV []rune

	fmt.Scan(&s)

	coda := insertCoda(s)
	codaS, codaC, codaV = spartisci(coda)

	switch c := os.Args[1]; c {
	case "S":
		codaS = extract(codaS)
	case "V":
		codaV = extract(codaV)
	case "C":
		codaC = extract(codaC)
	default:
		fmt.Println("Comando non riconosciuto")
	}

}
