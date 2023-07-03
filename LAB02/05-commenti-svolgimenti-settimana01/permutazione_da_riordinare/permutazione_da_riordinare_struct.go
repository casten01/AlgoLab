/* 
Il programma implementa una buona idea algoritmica, ma l'implementazione presenta difetti importanti:
- inizializzazione della slice di strutture
- invece di copiare tutti i campi delle strutture, bastava scambiare gli elementi della slice
*/

package main

import (
	"fmt"
)

type elemento struct {
	indice int
	nome   string
}

func stampaSlice(slice []elemento) {

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i].indice, slice[i].nome)
	}
}

func main() {

	var e elemento

	slice := make([]elemento, 0)
	e.indice = 6
	e.nome = "Francesco"
	slice = append(slice, e)

	e.indice = 1
	e.nome = "Andrea"
	slice = append(slice, e)

	e.indice = 5
	e.nome = "Elisa"
	slice = append(slice, e)

	e.indice = 2
	e.nome = "Beatrice"
	slice = append(slice, e)

	e.indice = 3
	e.nome = "Carlo"
	slice = append(slice, e)

	e.indice = 4
	e.nome = "Dino"
	slice = append(slice, e)

	e.indice = 7
	e.nome = "Giorgia"
	slice = append(slice, e)

	e.indice = 9
	e.nome = "Irene"
	slice = append(slice, e)

	e.indice = 8
	e.nome = "Henry"
	slice = append(slice, e)

	fmt.Println(slice)

	stampaSlice(slice)

	lungh := len(slice)

	for i := 0; i < len(slice); i++ { //posizione=(lunghezza-1)-indice elemento
		e.indice = slice[lungh-(slice[i].indice)].indice //e.indice=3
		e.nome = slice[lungh-(slice[i].indice)].nome
		slice[lungh-(slice[i].indice)].indice = slice[i].indice
		slice[lungh-(slice[i].indice)].nome = slice[i].nome //2=3
		slice[i].nome = e.nome                              //bisogna sostituire prima il nome e poi l'indice perch� altrimenti il valore dell'indice cambier�
		slice[i].indice = e.indice
	}

	stampaSlice(slice)

}


// REVISIONE:

//Dichiarazione e inizializzazione di una slice di strutture:
var slice = []elemento{
	{6, "Francesco"},
	{1, "Andrea"},
	{5, "Elisa"},
	...
}



//Le righe da 73 a 78 si possono scrivere molto più sinteticamente:
slice[i], slice[lengh-(slice[i].indice)] = slice[lingh-(slice[i].indice)] , slice[i]
