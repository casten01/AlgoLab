//DOCUMENTI HTML BEN FORMATI: scrivete un programma che legga una sequenza di tag e stabilisca se costituisce un documento html ben formato oppure no. Per farlo, potete usare una pila
/*
leggi un tag t;
se t è un tag di apertura
	inserisci t nella pila;
se t è un tag di chiusura
	estrai il tag t2 dalla cima della pila;
	se t e t2 non corrispondono
		il documento non ben formato;
*/
//Il ciclo andrà ripetuto finché non ci sono più tag da leggere. Alla fine, affinché il documento sia ben formato, la pila deve essere vuota (pensate a degli esempi di input in cui restano tag nella pila alla fine dell’esecuzione).
//errore con posizione nel caso documento non sia ben formato

package main

import (
	"fmt"
	"strings"
)

func pop(pila *[]string) string {
	//per copiare creo nuova pila, slicing e cambio puntatore
	var tmp []string
	var poppato string = (*pila)[len(*pila)-1]

	for i := 0; i < len(*pila)-1; i++ {
		tmp = append(tmp, (*pila)[i])
	}

	(*pila) = tmp

	return poppato
}

func push(pila *[]string, tag string) {
	*pila = append(*pila, tag)
}

func readTag(istruzioni string) []string {
	tags := strings.Split(istruzioni, " ")

	return tags
}

func benFormato(tags *[]string, pila *[]string) {

	for i := 0; i < len(*tags); i++ {
		var tmp string = (*tags)[i]
		var tmp2 string

		if tmp[1] != '/' {
			push(pila, (*tags)[i])
		} else {
			tmp2 = pop(pila)
			if tmp2[1] != tmp[2] {
				fmt.Println("Documento non ben formato")
				return
			}
		}
	}

}

func main() {
	var istruzioni string = "<a> <b> </b> <c> <d> </d>"
	var tags []string
	var pila []string
	tags = readTag(istruzioni)

	benFormato(&tags, &pila)

	fmt.Println(pila)
}
