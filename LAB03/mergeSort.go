//ALGORITMO PER FUSIONE: La funzione:
/*
1. divide il vettore in due sotto-vettori di dimensione circa uguale;
2. ordina il sotto-vettore di sinistra richiamando se stessa;
3. ordina il sotto-vettore di destra richiamando se stessa;
4. integra (merge) i due vettori in un vettore ordinato.
La base della ricorsione è, anche qui, data dai vettori di lunghezza 0 o 1, che sono sempre
ordinati.
La parte di integrazione (merge) di due vettore ordinati a1 e a2 funziona con un vettore di
supporto: si scorrono entrambi i vettori da sinistra a destra usando due indicatori i1 e i2 rispettivamente;
ad ogni passo si confronta a1[i1] con a2[i2] e si sceglie l’elemento più piccolo, lo si
copia nel vettore di supporto (nella prima posizione libera) e si incrementa l’indicatore relativo
ad esso. Quando i1 esce da a1 oppure i2 esce da a2, la parte rimanente dell’altro vettore viene
copiata nel vettore di supporto. Alla fine si copia il contenuto del vettore di supporto nel vettore
originale.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInteri() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	sequenza := strings.Split(text, " ")
	var seqInt []int
	for i := 0; i < len(sequenza); i++ {
		tmp, _ := strconv.Atoi(sequenza[i])
		seqInt = append(seqInt, tmp)
	}

	return seqInt
}

func mergeSort(seq []int, sx, dx int) {
	if sx < dx {
		m := (sx + dx) / 2
		mergeSort(seq, sx, m)
		mergeSort(seq, m+1, dx)
		merge(seq, sx, dx, m)
	}
	return
}

func merge(seq []int, sx, dx, m int) {
	i, j := sx, m+1

	var supp []int

	for i <= m && j <= dx {
		if seq[i] < seq[j] {
			supp = append(supp, seq[i])
			i++
		} else {
			supp = append(supp, seq[j])
			j++
		}
	}

	for i <= m {
		supp = append(supp, seq[i])
		i++
	}

	for j <= dx {
		supp = append(supp, seq[j])
		j++
	}

	for i := sx; i <= dx; i++ {
		seq[i] = supp[i-sx]
	}

}

func main() {
	seq := readInteri()
	n := len(seq)
	mergeSort(seq, 0, n-1)
	fmt.Println(seq)
}

/*
Gli algoritmi richiamati sopra possono essere implementati con stili diversi.
• Si può scegliere un’implementazione più “letterale” in cui l’algoritmo in pseudocodice è
tradotto letteralmente. Questo stile rende più evidenti le operazioni che si devono svolgere
e può aiutare a comprendere e valutare la complessità degli algoritmi.
• Si può preferire un uso più idiomatico del linguaggio, sfruttando ad esempio il doppio
assegamento, o notazioni sintattiche come a[low:high]. Questo stile richiede una maggiore
conoscenza del linguaggio di programmazione usato e può rendere i programmi più
sintetici o più leggibili (per chi conosce bene le specificità del linguaggio!); d’altra parte
potrebbe portare a sottovalutare i costi di certe operazioni “nascoste”.
*/
