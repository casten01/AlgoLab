/* 
DIFETTO 1 (algoritmo):
Il programma risolve il problema "letteralmente", cercando cioè le salite.
Basterebbe invece contare i punti di inizio delle salite, cioè i punti a V.

DIFETTO 2 (qualità del codice):
Anche l'implementazione presenta dei difetti. Nel programma si riconosce l'uso del piano per l'elaborazione
di due valori adiacenti in una serie. Il piano però non è ben implementato:
- avendo salvato i dati in un vettore, non serve usare le due variabili prev-curr, 
meglio usare gli indici i, i+1 
- l'uso delle variabili prev-curr è utile se non si vuole memorizzare la serie in un vettore

*/

package main

import "fmt"

func main(){
	sequenza := []int{9, 1, 3, 5, 2, 0, 8, 6}

	numeroSalite := 0

	salita := false
	numPrecedente := sequenza[0]
	for i:=1; i<len(sequenza); i++{
		numCorrente := sequenza[i]
		if numCorrente >= numPrecedente{
			if !salita{
				salita = true
				numeroSalite++
			}
		} else{
			salita = false
		}
		numPrecedente = numCorrente
	}

	fmt.Println(numeroSalite)
}


// versione rivista (difetto 2):
for i:=1; i<len(sequenza); i++{
	if sequenza[i] >= sequenza[i-1]{
		if !salita{
			salita = true
			numeroSalite++
		}
	} else{
		salita = false
	}
}