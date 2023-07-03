// SOTTOSTRINGHE a-b: Data una stringa di N caratteri nell’alfabeto a b c, stampare il numero di sottostringhe che iniziano con a e finiscono con b (tali sottostringhe possono sovrapporsi). Esempio: nella stringa 'ccbaacbabbcbab' il numero di sottostringhe a-b è 15. Notate che ciascuna delle prime due a (cioè, le due più a sinistra) appaiono ciascuna in 5 sottostringhe a-b).
//-> OSSERVAZIONE: Ogni b chiude tante sottostringhe a-b quante sono le a che la precedono.

package main

import (
	"fmt"
)

func readStringa() string {
	var stringa string
	fmt.Scan(&stringa)

	return stringa
}

func countAB() {
	parola := readStringa()
	var sottoStringhe int
	for i, c := range parola {
		if string(c) == "a" {
			for _, k := range parola[i:] {
				if string(k) == "b" {
					sottoStringhe += 1
				}
			}
		}
	}

	fmt.Println("Numero di sottostringhe:", sottoStringhe)
}

//VERSIONE LINEARE: conto le a prima di ogni b => sono il numero di sottostringhe ab
/*for i, ch := range sc.Text() {
	if ch == 'a' {
		countA++
	}
	if ch == 'b' {
		count += countA
	}
}
*/

func main() {
	countAB()
}
