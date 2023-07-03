//MASSIMO NUMERI di ZERO: Scrivere un programma massimo_n_zeri.go che legge da standard input una sequenza di interi terminata da -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi siano più numeri che contengono il massimo numero di 0, il programma stampa l'ultimo incontrato. Ad esempio, se la sequenza letta è 3040 145 80 1707 105002 78 1970 6005 -1 il programma stampa 105002.

package main

import (
	"fmt"
	"strconv"
)

func countZero(sequenza []int) map[string]int {
	cartina := make(map[string]int)

	for i := 0; i < len(sequenza); i++ {
		stringN := strconv.Itoa(sequenza[i])
		for _, v := range stringN {
			if v == '0' {
				cartina[stringN] += 1
			}
		}
	}
	return cartina
}

func maxZero(sequenza []int) {
	cartina := make(map[string]int)
	cartina = countZero(sequenza)

	var max string
	for k, v := range cartina {
		if v >= cartina[max] {
			max = k
		}
	}
	fmt.Println("Numero con massimo numero di zeri:	", max)
}

func main() {
	var sequenza []int
	var n int

	for {
		fmt.Scan(&n)
		if n == -1 {
			break
		} else {
			sequenza = append(sequenza, n)
		}
	}
	maxZero(sequenza)
}
