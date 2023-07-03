// SOTTOSEQUENZE CRESCENTI: Data una lunga sequenza di N interi distinti che rappresentano le altezze di una catena montuosa, stampare il numero di salite che vanno da sinistra a destra (una salita è una sequenza crescente di 2 o più interi, che partono da un punto di minimo e arrivano in un punto di massimo). Esempio: nella sequenza 9 1 3 5 2 0 8 6 ci sono due salite: 1 3 5 e 0 8 (1 3 e 3 5 non sono salite perché la prima non finisce in un punto di massimo e la seconda non inizia in un punto di minimo).

//-> Il numero di salite è uguale al numero di inizi di tali salite. E un inizio si ha quando ci sono 3 interi successivi a forma di V: grande-piccolo-grande. NB: attenzione ai bordi! OSSERVAZIONE DEMMERDA

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

func numeroSalite(sequenza []int) int {

	var stessaSalita bool = false
	var numSalite int

	for i := 1; i < len(sequenza); i++ {
		if sequenza[i] > sequenza[i-1] && stessaSalita != true {
			stessaSalita = true
			numSalite += 1
		} else {
			stessaSalita = false
		}
	}

	return numSalite
}

func main() {
	seq := readInteri()
	num := numeroSalite(seq)
	fmt.Println(num)
}
