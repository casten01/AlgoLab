//1.3 IMPLEMENTAZIONE: algoritmo che determina se H possa raggiungere la casa di S: restituisce il numero di gallerie percorse per arrivarci o -1 altrimenti. Regola: prendere galleria di luminosità minore
/*
Formato di input:
- prima riga: N M H S
	-> con N numero di svincoli, M numero di gallerie, H indice svincolo di partenza, S di arrivo
- altre righe: A B L
	-> A B svincoli collegati da galleria, L luminosità della galleria
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type svincolo struct {
	visitato bool
	gallerie map[int]int //->chiave: indice dello svincolo raggiungibile, valore: luminosità della galleria
}
type grafo []svincolo //posizione indica indice di quello svincolo

// legge il grafo secondo il formato di input descritto
func readGrafo() (g grafo, H int, S int) {
	f, _ := os.Open("inputVampiri.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), " ")
		if len(info) > 3 {
			N, _ := strconv.Atoi(info[0])
			//M, _ := strconv.Atoi(info[1])
			H, _ = strconv.Atoi(info[2])
			S, _ = strconv.Atoi(info[3])
			g = make([]svincolo, N)
			for k := range g {
				g[k].gallerie = make(map[int]int)
				g[k].visitato = false
			}
			continue
		}
		A, _ := strconv.Atoi(info[0])
		B, _ := strconv.Atoi(info[1])
		L, _ := strconv.Atoi(info[2])
		g[A-1].gallerie[B] = L
		g[B-1].gallerie[A] = L

	}

	return g, H, S

}

// Stampa il grafo; attenzione! La slice []svincolo di g ha gli indici che partono da 0 e corrispondono ai numeri dei nodi-1
func printGrafo(g grafo) {
	for k, v := range g {
		fmt.Printf("%d collegato con: ", k+1)
		for k2, b := range v.gallerie {
			fmt.Printf("[%d] ->lum: %d, ", k2, b)
		}
		fmt.Println()
	}
}

// Algoritmo: parto dal nodo di partenza H, lo marco come visitato e controllo possibili archi uscenti e scelgo quello di luminosità minore, marco il nodo a cui arrivo come visitato. Se l'arco che devo scegliere mi porta a un nodo che è già stato visitato mi fermo(sono in un ciclo), altrimenti continuo finchè non raggiungo il nodo S
func trovaAmica(g grafo, H int, S int) (res int) {
	var l_min int = 1000 //standard max di luminosità
	var k_min int
	var next int = H //parto da casa di Harmony
	//finchè non rivisito lo stesso nodo, in quel caso sono in un ciclo
	for g[next-1].visitato != true {
		g[next-1].visitato = true
		if next == S {
			return res
		}

		//le chiavi della mappa corrispondo direttamente al numero dello svincolo
		for k, v := range g[next-1].gallerie {
			if v < l_min {
				l_min = v
				k_min = k
			}
		}
		l_min = 1000
		res++
		next = k_min //svincolo in cui sono andato

	}

	return -1
}

func main() {
	g, H, S := readGrafo()
	fmt.Printf("Da %d a %d \n", H, S)
	printGrafo(g)
	res := trovaAmica(g, H, S)
	fmt.Println(res)
}
