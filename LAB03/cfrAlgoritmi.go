// CONFRONTO DI ALGORITMI: si considerino i due algoritmi AlgoX and AlgoY riportati sotto e si svolgano i seguenti punti.
package main

import "fmt"

func main() {
	var table = []int{-9, -1, 0, 13, 14, 14, 12, 29, 31, 24, 36, 36, 44, 44, 8}

	fmt.Println(algoY(table, 14))
}

// 1. Si tracci algoX (cioè: si simuli l’esecuzione tenendo traccia dello stato del programma) per table indicata qui sotto e x pari a 14. Si elenchino in particolare tutti i valori della variabile i durante l’esecuzione.

/* x = 14
i=5
*/
// 2. Si tracci AlgoY (cioè: si simuli l’esecuzione tenendo traccia dello stato del programma) per table indicata qui sotto e x pari a 14. Si scrivano in particolare i valori della variabili low, mid e high subito dopo l’esecuzione della riga 6 (ogni volta che viene eseguita).
/* x=14 low=0, high=0
1] m7 l0 h6
2] m3 l4 h6
3] m5 true
*/
// 0  1 2  3  4  5  6  7  8  9 10 11 12 13 14
//-9 -1 0 13 14 14 12 29 31 24 36 36 44 44 8

// 3. Si indichi quali delle seguenti affermazioni sono vere per AlgoX e/o per AlgoY. Per ciascuna affermazione si usi una delle cinque opzioni seguenti: X (= l’affermazione è corretta solo per AlgoX), Y (= l’affermazione è corretta solo per AlgoY), X & Y (= l’affermazione è corretta sia per per AlgoX che per AlgoY), no (= l’affermazione non è corretta né per AlgoX né per AlgoY), non so (= non so la risposta).
// a) L’algoritmo esamina gli elementi partendo dall’indice minimo al massimo. [X]
// b) L’algoritmo cerca l’elemento massimo di table. [NO]
// c) Come secondo valore, l’algoritmo restituisce sempre l’indice minore per l’elemento x, se questo è contenuto in table. [X]
// d) L’algoritmo ordina gli elementi in table. [NO]
// e) L’algoritmo restituisce tutti gli indici in cui si trova l’elemento x. [NO]
// f) L’algoritmo esamina tutti gli elementi di table. [NO]
// g) L’algoritmo è corretto solo se table è ordinato. [Y]
// h) L’algoritmo alla fine restituisce sempre false, -1. [NO]

// 4. Si argomenti se la seguente affermazione è vera o falsa: “AlgoX è più efficiente di AlgoY per cercare un singolo elemento in un vettore”.
//dipende dove è l'elemento

// 5. Quali nomi più esplicativi si potrebbero dare agli algoritmi AlgoX e AlgoY?
//ricerca lineare e ricerca binaria

// 6. È facile modificare AlgoX in modo che restituisca tutti gli indici in cui si trova x? Cambia la complessità dell’algoritmo?
//sì cambia, basta salvare gli indici che vengono restituiti, ma bisogna scansionare tutto il vettore per forza

// 7. È facile modificare AlgoY in modo che restituisca tutti gli indici in cui si trova x? Cambia la complessità dell’algoritmo?
//

// Listing 1: AlgoX
func algoX(table []int, x int) (bool, int) {
	for i, el := range table {
		if el == x {
			return true, i
		}
	}
	return false, -1
}

// Listing 2: AlgoY
func algoY(table []int, x int) (bool, int) {
	low, high := 0, len(table)-1

	for low <= high {
		mid := (low + high) / 2
		if table[mid] < x {
			low = mid + 1
		} else if table[mid] > x {
			high = mid - 1
		} else {
			return true, mid
		}
	}

	return false, -1
}
