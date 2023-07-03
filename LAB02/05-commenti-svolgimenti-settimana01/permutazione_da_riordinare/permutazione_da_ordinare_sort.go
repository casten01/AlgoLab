/* DIFETTO: Si risolve il problema riordinando la sequenza,
senza sfruttare il fatto che si tratta di una permutazione
(in particolare, questo implica che sappiamo già, a priori,
in quale posizione finiranno i vari elementi della sequenza!)
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type persona struct {
	id   int
	nome string
}

// Restituisce un puntatore ad una nuova persona
func nuovaPersona(id int, nome string) *persona {
	p := persona{id: id, nome: nome}
	return &p
}

// Costruisce un array di puntatori a persone basandosi sul contenuto
// di un array di stringhe in formato "id nome"
func personeFromArrayOfStrings(stringPersons []string) []*persona {
	var persone []*persona

	for _, person := range stringPersons {
		splitted := strings.Split(person, " ")
		id, err := strconv.Atoi(splitted[0])
		if err == nil {
			persona := *nuovaPersona(id, splitted[1])
			persone = append(persone, &persona)
		}
	}
	return persone
}

// Scambia il puntatore in posizione a con il puntatore in posizione b
func scambia(arr []*persona, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

// Stampa un array di puntatori a persona
func print(arr []*persona) {
	fmt.Println("Persone:")
	for _, persona := range arr {
		fmt.Println(*persona)
	}
}

// bubbleSort
func bubbleSort(persons []*persona) {
	for k := 1; k < len(persons); k++ {
		for i := 1; i < len(persons)-k+1; i++ {
			if persons[i-1].id > persons[i].id {
				scambia(persons, i-1, i)
			}
		}
	}
}

// bubbleSortReverse
func bubbleSortReverse(persons []*persona) {
	for k := 1; k < len(persons); k++ {
		for i := 1; i < len(persons)-k+1; i++ {
			if persons[i-1].id < persons[i].id {
				scambia(persons, i-1, i)
			}
		}
	}
}

// Main
func PermutazioneDaOrdinare() {
	personeString := []string{
		"6 Francesco",
		"1 Andrea",
		"5 Elisa",
		"2 Beatrice",
		"3 Carlo",
		"4 Dino",
		"7 Giorgia",
		"9 Irene",
		"8 Henry"}
	persone := personeFromArrayOfStrings(personeString)
	print(persone)
	bubbleSort(persone)
	print(persone)
	persone = personeFromArrayOfStrings(personeString)
	bubbleSort(persone)
	print(persone)
	persone = personeFromArrayOfStrings(personeString)
	bubbleSortReverse(persone)
	print(persone)

}
