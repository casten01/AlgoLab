//PRIME VOCALI: Si vuole scrivere un programma che legga una sequenza di stringhe e stampi, per ogni stringa, la sua prima vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole). Nel caso in cui una stringa non contenga vocali il programma stampa “la parola non contiene vocali”.

package main

import (
	"fmt"
)

func firstVocale(word string) string {
	var vocali = []string{"a", "e", "i", "o", "u"}

	for _, v := range word {
		for i := 0; i < 5; i++ {
			if string(v) == vocali[i] {
				return string(v)
			}
		}
	}

	return " "
}
func main() {
	cartinaLunga := make(map[string]string)
	var word string
	for i := 0; i < 5; i++ {
		fmt.Scan(&word)
		cartinaLunga[word] = firstVocale(word)
	}

	for k, v := range cartinaLunga {
		if cartinaLunga[k] != " " {
			fmt.Printf("%s -> prima vocale: %s \n", k, v)
		} else {
			fmt.Printf("%s -> non contiene vocali \n", k)
		}
	}
}
