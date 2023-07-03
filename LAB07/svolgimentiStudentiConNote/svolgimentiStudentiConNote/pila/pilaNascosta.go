package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	var numeri []int

	lenSlice = len(numeri) - 1

	switch value {
	case "+":
		oper = numeri[lenSlice] + numeri[lenSlice-1]
		numeri = numeri[0:(lenSlice - 1)]
		numeri = append(numeri, oper)
	
	...
	default:
		num, _ := strconv.Atoi(value)
		numeri = append(numeri, num)
	}
	...
	fmt.Println("Il risultato è: ", numeri[lenSlice-1])
}



/* Note:

In questo programma, l'uso della pila, implementata con una slice, è implicito. 

Di fatto la slice di interi "numeri" è usata come una pila:
- si inseriscono elementi solo in fondo, con append
- si tolgono elementi solo dal fondo, con numeri[lenSlice]

Tuttavia non c'è alcuna indicazione visibile del fatto che "numeri" sia una pila:
- il nome della variabile non fa riferimento alla pila
- non ci sono funzioni con i nomi tipici (push, pop)

Si tratta di una soluzione ad hoc, veloce da scrivere, poco adatta al riuso del codice.

*/