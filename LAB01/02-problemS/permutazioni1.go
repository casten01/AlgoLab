//PERMUTAZIONI DA RIORDINARE: Si consideri un vettore di N strutture, ciascuna composta da un nome e da una chiave identificativa da 1 a N. Le chiavi sono uniche, cioè non sono ripetute. Progettate un algoritmo che riordini il vettore in questo modo: nella posizione 0 c’è l’elemento di chiave N, nella posizione 1 c’è l’elemento di chiave N􀀀1 e così via fino alla posizione N-1 che dovrà contenere l’elemento di chiave 1.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Persona struct {
	nome string
	id   int
}

const N int = 8

func createLista() [N]Persona {
	var persone [N]Persona
	var i int

	// open file
	f, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	//defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		persone[i].nome = scanner.Text()
		persone[i].id = i + 1
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return persone
}

func ordinaListaRev(persone [N]Persona) [N]Persona {
	var ind_max int
	var tmp Persona

	for i := 0; i < N-1; i++ {
		ind_max = i
		for j := i; j < N; j++ {
			if persone[j].id >= persone[ind_max].id {
				ind_max = j
			}
		}
		tmp = persone[i]
		persone[i] = persone[ind_max]
		persone[ind_max] = tmp

	}

	return persone
}

func main() {
	persone := createLista()
	fmt.Println(persone)
	personeOrd := ordinaListaRev(persone)
	fmt.Println(personeOrd)

}
