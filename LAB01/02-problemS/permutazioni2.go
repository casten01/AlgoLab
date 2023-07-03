//RACCOGLIERE PERMUTAZIONI: Data una permutazione di 1..N, vogliamo raccogliere i numeri in ordine crescente cominciando ad analizzarli da sinistra. Scrivete un programma che stabilisce quante volte avremo bisogno di tornare verso sinistra?
//Esempio: Nella permutazione 4 5 1 3 6 2 l’output è 2, poiché 1 si trova andando sempre verso destra, poi si prosegue verso destra per raccogliere 2, ma per raccogliere 3 bisogna tornare indietro verso sinistra; bisogna tornare ancora indietro per raccogliere 4, dopodiché 5 e 6 si trovano in ordine proseguendo verso destra.

package main

import "fmt"

func findNext(sequenza []int) int {
	var tmp, tmp2, mv_sx int

	for i := 0; i < len(sequenza)-1; i++ {
		tmp = 0
		tmp2 = 0

		for j := 0; j < len(sequenza); j++ {
			if sequenza[j] == i+1 {
				tmp = j
			}
			if sequenza[j] == i+2 {
				tmp2 = j
			}
		}
		if tmp2 < tmp {
			mv_sx += 1
		}
	}

	return mv_sx
}

//VERSIONE TEMPO LINEARE: ha complessità lineare
func raccogli(v []int) int {
	count := 0
	occ := make(map[int]bool)
	for _, el := range v {
		occ[el] = true
		if occ[el+1] {
			count++
		}
	}
	fmt.Printf("Servono %d inversioni\n", count)
	return count
}

func main() {
	var sequenza []int
	var num int

	for i := 0; i < 6; i++ {
		fmt.Scan(&num)
		sequenza = append(sequenza, num)
	}

	fmt.Println(findNext(sequenza))
}
