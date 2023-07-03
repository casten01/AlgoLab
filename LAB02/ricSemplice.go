//1) FUNZIONE RICORSIVA SEMPLICE: Considerate la seguente funzione, che deve restituire il prodotto dei suoi due argomenti Cosa si deve scrivere, nel caso base, al posto di AAAAA e di BBBBB?

package main

import "fmt"

func multiply(x, y int) int {
	if y == 1 {
		return x
	} else {
		return x + multiply(x, y-1)
	}
}

func main() {
	fmt.Println(multiply(5, 3))
}
