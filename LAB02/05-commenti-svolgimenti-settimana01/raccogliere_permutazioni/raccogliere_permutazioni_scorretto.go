/* DIFETTO:
Il problema viene semplificato eccessivamente:
si contano tutti i "cali" (punti in cui si ha un numero seguito da uno inferiore),
ma questi non sono indicativi del problema.

Ad esempio, nella sequenza 5 1 2 6 3 ci sono due cali (5 1, 6 3)
ma solo un'inversione (dopo 3).

Dunque il programma ha complessità lineare, ma NON E' CORRETTO,
poiché l'output prodotto su certe istanza è sbagliato
*/

package main

import "fmt"

func main() {
	var check int
	slice := []int{4, 5, 1, 3, 6, 2}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			check++
		}
	}
	fmt.Println(check)
}
