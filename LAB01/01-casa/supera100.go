/*SUPERA 100: Scrivere un programma supera100.go che legge da standard input una sequenza di interi positivi
terminata da -1 e stampa il primo numero che supera 100, se presente; altrimenti stampa
“nessun numero maggiore di 100”.*/

package main

import "fmt"

func main() {

	sequenza := make([]int, 1)
	var i int = 0
	var n int = 0

	for i > -1 {
		fmt.Scan(&n)
		if n != -1 {
			sequenza = append(sequenza, n)
			i++
			continue
		}
		break
	}

	for j := 0; j <= i; j++ {
		if sequenza[j] > 100 {
			fmt.Println(sequenza[j])
			break
		} else if j == i {
			fmt.Println("nessun numero maggiore di 100")
		} else {
			continue
		}
	}
}
