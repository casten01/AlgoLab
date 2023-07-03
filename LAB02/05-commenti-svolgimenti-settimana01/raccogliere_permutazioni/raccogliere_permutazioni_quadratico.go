/* Il programma considera gli interi k da 1 a N.
Si usa la variabile pos per tenere traccia della posizione del k precedente (all'inizio pos è zero)
Per ciascun k, cerca k nella sequenza, trova la sua posizione i
e verifica se i è successivo ("a destra di") a pos. In questo caso,
significa che è necessaria un'inversione.const

DIFETTTO: anche se il programma è corretto,
per ogni k si deve scorrere l'intera sequenza.
La complessità quindi è quadratica, ma il problema si potrebbe risolvere linearmente.

Anche aggiungendo un break a riga 36,
nel caso peggiore il costo rimane quadratico
(es: 5 4 3 2 1)


*/

package main

import (
	"fmt"
)

const N = 6

func main() {
	sequenza := make([]int, N)
	var pos, sx int

	for i := 0; i < N; i++ {
		fmt.Scan(&sequenza[i])
	}

	for k := 1; k < N; k++ {
		for i := 0; i < len(sequenza); i++ {
			if sequenza[i] == k {
				if i < pos {
					sx++
				}
				pos = i
			}
		}
	}
	fmt.Println("output:", sx)
}
