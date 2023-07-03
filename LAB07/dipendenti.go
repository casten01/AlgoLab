//DIPENDENTI: lavoro organizzato in maniera gerarchica; ogni dipendente (tranne quelli di massimo livello) ha un supervisore di cui è detto subordinato (possono essere 0, 1, +1) [subordinati sono solo quelli direttamente collegati]
/*
(a) Dato un certo dipendente, stampare l’elenco dei suoi subordinati
(b) Contare quanti sono i dipendenti che non hanno alcun subordinato.
(c) Dato un certo dipendente, individuare chi è il suo supervisore.
(d) Dato un certo dipendente, stampare la lista dei dipendenti che si trovano sopra di lui gerarchicamente,
partendo dal suo supervisore e risalendo la gerarchia fino a un dipendente di
massimo livello.
(e) Stampare l’elenco di tutti i dipendenti –non importa l’ordine–, indicando per ciascuno chi è
il suo supervisore (tranne nel caso di dipendenti di massimo livello).
*/

//da riscrivere con puntatore alla mappa, la passo come argomento impiegati *map[string]string e accedo con (*impiegati)[dip] oppure range *impiegati
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//mappa padri: chiave padre -> figli
//albero con nodi formati da una stringa nome + puntatore padre + slice di puntatori ai figli subordinati

//input: dipendente+mappa; output: stampa elenco suoi subordinati
func stampaSubordinati(dip string, impiegati map[string]string) {
	fmt.Printf("I subordinati di %s sono:", dip)
	for k, v := range impiegati {
		if v == dip {
			fmt.Printf(" %s", k)
		}
	}
	fmt.Println()
}

//input: mappa; output: numero dipendenti senza subordinati
func quantiSenzaSubordinati(impiegati map[string]string) int {
	var counter int = len(impiegati)

	for k := range impiegati {
		for _, v := range impiegati {
			if v == k {
				counter--
				break
			}
		}
	}
	return counter
}

//input: dipendente + mappa; output: supervisore del dipendente
func supervisore(dip string, impiegati map[string]string) string {
	var super string

	for k, v := range impiegati { //posso accedere direttamente: super = impiegati[dip], non avrei controllo su quelli che non hanno supervisore
		if k == dip {
			super = v
			return super
		}
	}
	super = "non ha supervisori"
	return super
}

//input: dip+mappa; output: gerarchia dei suoi supervisori
func stampaImpiegatiSopra(dip string, impiegati map[string]string) {
	var super string

	for k, v := range impiegati {
		if k == dip {
			super = v
			fmt.Printf("%s ", super)
			stampaImpiegatiSopra(super, impiegati)
		}
	}
	return
}

//input: mappa; output: elenco di tutti i dipendenti con supervisore
func stampaImpiegatiConSupervisore(impiegati map[string]string) {
	for k, v := range impiegati {
		fmt.Print("impiegato: ", k, " -> supervisore : ", v)
		fmt.Println()
	}
}

//legge gerarchia dipendenti da file input.txt
func readFromFile() []string {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var istruzioni []string
	for scanner.Scan() {
		istruzioni = append(istruzioni, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return istruzioni
}

//riempie mappa con coppie subordinato - supervisore
func fillImpiegati(impiegati map[string]string, gerarchie []string) map[string]string {
	for i := 0; i < len(gerarchie); i++ {
		str := strings.Split(gerarchie[i], "-")

		impiegati[str[0]] = str[1]
	}

	return impiegati
}

func main() {
	impiegati := make(map[string]string)
	gerarchie := readFromFile()
	impiegati = fillImpiegati(impiegati, gerarchie)

	// fmt.Println(impiegati)
	// stampaSubordinati("anna", impiegati)
	// fmt.Println("Numero dipendenti senza subordinati: ", quantiSenzaSubordinati(impiegati))
	// fmt.Println("Il supervisore di anna:", supervisore("anna", impiegati))
	// stampaImpiegatiSopra("irene", impiegati)
	// stampaImpiegatiConSupervisore(impiegati)

	fmt.Println(impiegati["anna"])

}
