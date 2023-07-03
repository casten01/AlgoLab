/*ELABORAZIONE DI UNA SERIE DI INTERI:*/

package main

import (
	"fmt"
	"strings"
)

/*
1) Scrivere una funzione quanteConA(parole []string) int che, data una slice di stringhe, restituisca quante stringhe contengono il carattere ‘a’.
*/

func quanteConA(parole []string) int {
	var counter int = 0
	for i := 0; i < len(parole); i++ {
		if strings.Contains(parole[i], "a") {
			counter += 1
		}
	}

	return counter
}

/*
2) Scrivere una funzione primaConA(parole []string) string che, data una slice di
stringhe, restituisca la prima parola che contiene il carattere 'a', o la stringa vuota.
*/

func primaConA(parole []string) {
	var primaA string
	for i := 0; i < len(parole); i++ {
		if strings.Contains(parole[i], "a") {
			primaA = parole[i]
			fmt.Println(primaA)
			return
		}
	}
	fmt.Println(primaA)
}

/*
3) Scrivere una funzione piuCorta(parole []string) int che, data una slice di stringhe,
restituisca la lunghezza della stringa più corta in termini di byte.
*/

func piuCorta(parole []string) int {
	var lenCorta int = len(parole[0])
	for i := 1; i < len(parole); i++ {
		if len(parole[i]) < lenCorta {
			lenCorta = len(parole[i])
		}
	}

	return lenCorta
}

func main() {

	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	primaConA(parole)
	fmt.Println(piuCorta(parole))
}
