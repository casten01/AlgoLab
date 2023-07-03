//MAPPA DELLE ORBITE - Advent Of Code 2019 day 6
//Alberi - implementazione con tabella dei padri: ci interessa relazione tra padre-figlio => per ciascun elemento si memorizza il padre (tabella dei padri->in Go uso una mappa)

// mappa delle orbite, ogni oggetto dello spazio, tranne il COM, ruota attorno a esattamente un altro oggetto,     AAA)BBB significa che BBB orbita attorno ad AAA

//PARTE 1: Se A orbita attorno a B e B orbita attorno a C, si dice che A orbita indirettamente attorno a C. Devi calcolare il numero di tutte le orbite dirette e indirette presenti nella mappa. Qual è il numero totale di orbite dirette e indirette

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFileAndMapping() map[string]string {
	// open file
	f, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var orbite = make(map[string]string) //dichiaro mappa vuoto
	var pianeta1, pianeta2 string

	//popolo la mappa:
	for scanner.Scan() {
		testo := scanner.Text()
		pianeti := strings.Split(testo, ")")
		pianeta1 = pianeti[0]
		pianeta2 = pianeti[1]

		orbite[pianeta2] = pianeta1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return orbite
}

func countOrbite(orbite map[string]string) int {
	var counter int

	//per ogni chiave della mappa risalgo alla radice e conto quanti nodi attraversa; sommo sempre allo stesso contatore trovando il numero di orbite indirette
	for k, _ := range orbite {
		var tmp string = orbite[k]
		for tmp != "COM" {
			counter++
			nonno := orbite[tmp]
			tmp = nonno
		}
	}

	//aggiungo orbite dirette, una per ogni coppia della mappa
	counter += len(orbite)

	return counter
}

func main() {
	orbite := readFileAndMapping()
	numOrbiteTot := countOrbite(orbite)
	fmt.Println(numOrbiteTot)
}

//1) Modellate la situazione con un albero: cosa rappresentano i nodi dell'albero? cosa rappresenta la relazione padre-figlio?

//2) Cosa sono le orbite dirette? E le orbite indirette? Come può essere descritto, in termini di alberi, il numero di orbite indirette di un oggetto?
