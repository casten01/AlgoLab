package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func controlla(s string) {
	pila := make([]string, 0)
	tags := strings.Fields(s)
	for pos, tag := range tags {
		if tag[1] != '/' {
			pila = append(pila, tag)
		} else {
			if len(pila) == 0 {
				fmt.Println("Errore in pos:", pos+1)
				return
			}
			tag2 := pila[len(pila)-1]
			if tag[2] != tag2[1] {
				fmt.Println("Errore in pos:", pos+1)
				return
			}
			pila = pila[:len(pila)-1]
		}
	}
	if len(pila) != 0 {
		fmt.Println("Sono rimasti aperti i seguenti tag:", pila)
	} else {
		fmt.Println("Stringa ben formata")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	controlla(scanner.Text())
}

/* Note:

In questo programma, l'implementazione della pila è implicita,
poiché non vi sono funzioni con i nomi tipici (push, pop)

C'è però un rigerimento esplicito alla pila in riga 11, dove si dichiara una variabile "pila"
C'è quindi una manifestazione esplicita dell'intenzione di usare la struttura dati pila, anche senza implementarla come struttura di dati *astratta*.

Di fatto la pila è implementata (implicitamente) come slice di stringhe
- pila:=make([]string,0) (riga 11)
- si inseriscono elementi solo in fondo, con append
- si tolgono elementi solo dal fondo, con numeri[lenSlice]

Si tratta di una soluzione ad hoc, veloce da scrivere, poco adatta al riuso del codice.

*/
