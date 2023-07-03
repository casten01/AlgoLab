/*ELABORAZIONE DI UNA SERIE DI INTERI:*/
package main

import "fmt"

/*
1) Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro
una slice di interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il
prodotto. Ad esempio, se la slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà
restituire il numero 96 (pari al prodotto di 12 per 8).
*/
func stranoProdotto(numeri []int) int {
	var prodotto int = 1
	for i := 0; i < len(numeri); i++ {
		if numeri[i] > 7 && numeri[i]%2 == 0 {
			prodotto *= numeri[i]
		}
	}
	return prodotto
}

/*
2) Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice
di interi, stampi, per ciascun numero, se è pari o dispari.
*/

func pariDispari(numeri []int) {
	for i := 0; i < len(numeri); i++ {
		if numeri[i]%2 == 0 {
			fmt.Printf("%d e' pari\n", numeri[i])
		} else {
			fmt.Printf("%d e' dispari\n", numeri[i])
		}
	}
}

/*
3) Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi,
restituisca il più piccolo numero dispari (la slice può contenere sia numeri positivi che
negativi).
*/

func minDispari(numeri []int) int {
	var minDisp int = numeri[0]
	for i := 1; i < len(numeri); i++ {
		if numeri[i]%2 != 0 && numeri[i] < minDisp {
			minDisp = numeri[i]
		}
	}
	return minDisp
}

func main() {
	const N = 10
	numeri := make([]int, 10)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println(minDispari(numeri))
	//pariDispari(numeri)
}
