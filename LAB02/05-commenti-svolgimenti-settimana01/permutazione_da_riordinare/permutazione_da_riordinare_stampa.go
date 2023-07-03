/* DIFETTI:
- non fa quanto richiesto (stampa in ordine, invece di riordinare il vettore)
- complessità quadratica: per ogni k da 1 a N, scorre tutta la sequenza per cercare k
*/

package main

import (
	"fmt"
)

type Persona struct {
	nome   string
	chiave int
}

const N = 9

func main() {

	elenco := make([]Persona, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&elenco[i].chiave, &elenco[i].nome)
	}

	for k := N; k >= 1; k-- {
		for i := 0; i < len(elenco); i++ {
			if elenco[i].chiave == k {
				fmt.Println(elenco[i].chiave, elenco[i].nome)
			}
		}

	}

}
