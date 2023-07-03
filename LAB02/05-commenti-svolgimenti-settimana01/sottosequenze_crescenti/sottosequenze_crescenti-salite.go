/* Il programma risulta più contorto del dovuto, poiché risolve il problema "letteralmente",
cercando cioè le salite.
Basterebbe invece contare i punti di inizio delle salite, cioè i punti a V.
*/

package main

import "fmt"

func main() {
	array := []int{9, 1, 3, 5, 2, 0, 8, 6, 1, 2, 1, 2, 3, 4, 5, 6}
	fmt.Println(TrovaSalite(array))
}

func TrovaSalite(arr []int) int {
	steps := 0
	num_salite := 0

	for i := 0; i < len(arr); i++ {

		//Se sono ad un minimo o all'inizio
		if i == 0 || arr[i] < arr[i-1] { // DIFETTO: meglio gestire il caso i = 0 fuori dal ciclo
			// in ogni caso, quando i = 0, steps vale 0 quindi la condizione non è mai verificata
			//Se ho fatto almeno uno step segno la salita
			if steps > 0 {
				num_salite++
			}
			//Azzero gli steps
			steps = 0
			//Se sto facendo uno step in salita
		} else if arr[i] >= arr[i-1] {
			steps++
			//Se è l'ultimo step incremento il numero di salite
			if i == len(arr)-1 {
				num_salite++
			}
		}
	}
	return num_salite
}
