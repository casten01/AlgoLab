// ELETTRICITA': Vogliamo leggere una sequenza di N interi (almeno 3), che descrivono il consumo di elettricità in N giorni, e stampare i giorni in cui il consumo è stato inferiore rispetto sia al giorno prima che al giorno dopo. I giorni sono numerati a partire da 1.

package main

import "fmt"

func main() {
	const N int = 5
	var cons_day int
	var consumi [N]int

	for i := 0; i < N; i++ {
		fmt.Scan(&cons_day)
		consumi[i] = cons_day
	}

	for i := 1; i < N-1; i++ {
		if consumi[i] < consumi[i-1] && consumi[i] < consumi[i+1] {
			fmt.Printf("Giorno %d ° con consumo -> %d \n", i, consumi[i])
		}
	}
}
