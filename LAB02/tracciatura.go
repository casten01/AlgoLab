//2) TRACCIATURA DI FUNZIONE RICORSIVA: Considerate la seguente funzione, che deve calcolare il valore massimo contenuto nel vettore numbers. La funzione max calcola il massimo tra i suoi due argomenti.

package main

import "fmt"

func largest(numbers []int) int {
	n := len(numbers)

	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//1. Come deve essere completato il caso base? [ 1, 2, 5, 7, -2, 10, 9, 21, 3, 8 ]

//2. Durante l’esecuzione della chiamata largest(numbers), considerate la chiamata ricorsiva che termina per prima. Qual è l’argomento passato in questa chiamata e quale valore restituisce la funzione?

//-> viene passato il vettore [1,2] e restituisce 1

//3. Con quali argomenti viene eseguita per la prima volta la funzione max

//-> tra 2 e 1 e restituisce 2

//4. E con quali argomenti viene eseguita l’ultima volta la funzione max?

//-> tra 21 e 8

func main() {
	var numbers = [10]int{1, 50, 30, 7, -2, 10, 9, 21, 3, 8}

	fmt.Println(largest(numbers[:]))
}
