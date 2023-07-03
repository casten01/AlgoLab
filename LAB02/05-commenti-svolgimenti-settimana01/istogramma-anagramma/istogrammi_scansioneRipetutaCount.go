/* Si controllano solo parole della stessa lunghezza.
Si scandisce la prima parola, e per ogni lettera,
si fa una ricerca della lettera nella seconda parola.
Se la si trova, si aumenta il contatore count (alla fine, count deve essere pari alla lunghezza di s1,
che significa che ogni lettera di s2 è stata trovata) --- meglio una flag booleana?
Per gestire le lettere ripetute (devono essere nello stesso numero nelle due parole),
si cancella dalla seconda parola la lettera trovata.

DIFETTO (complessità):
la seconda parola viene scandita tante volte quanto è lunga la prima parola.
La complessità nel caso peggiore è dell'ordine di grandezza del prodotto delle due lunghezze n1*n2.

Usando una struttura di supporto (mappa, o slice) basterebb scansionare ciascuna stringa una sola volta
(complessità n1 + n2). Questo comporta un aumento dello spazio occupato, ma di dimensione costante
(la mappa ha tanti elementi quante le lettere dell'alfabeto)

DIFETTO di funzionamento: alla fine del programma, la seconda parola è stata svuotata.
*/

package main

import (
	"fmt"
)

func main() {
	var counter int
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	if len(s1) != len(s2) {
		fmt.Println("Le due parole non sono anagrammi")
		return
	}

	for i := 0; i < len(s1); i++ {
		//fmt.Println(s1, s2) // per mostrare come cambia s2
		c := s1[i]
		for k := 0; k < len(s2); k++ {
			if s2[k] == c {
				s2 = s2[:k] + s2[k+1:]
				counter++
				break
			}
		}
	}

	if counter == len(s1) {
		fmt.Println("Le due parole sono anagrammi")
	} else {
		fmt.Println("Le due parole non sono anagrammi")
	}

}
