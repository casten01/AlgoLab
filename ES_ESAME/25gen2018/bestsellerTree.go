//ALBERO BINARIO DI RICERCA PER BESTSELLER: vedi appunti ipad

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	titolo  string
	autore  string
	editore string
	data    string
	genere  string
}

type node struct {
	libro item
	dx    *node
	sx    *node
}

type tree struct {
	root *node
}

// Crea un nodo dell'albero
func newNode(val item) *node {
	return &node{val, nil, nil}
}

// Cerca dove attaccare nodo
func bist_searchparent2(r **node, n item) *node {
	p := *r
	var parent *node = nil

	for p != nil {
		res := cmp(n.data, p.libro.data)
		if !res {
			parent = p
			p = p.sx
		} else {
			parent = p
			p = p.dx
		}
	}
	return parent
}

// Inserire un nuovo elemento
func bist_insert(r *node, n item) string {
	if r == nil {
		return "L'albero è vuoto"
	}
	res := bist_searchparent2(&r, n)

	if !cmp(n.data, res.libro.data) {
		(*res).sx = newNode(n)
	} else {
		(*res).dx = newNode(n)
	}
	return "Elemento inserito correttamente"
}

// funzione che trasforma due stringhe nel formato mese/giorno/anno in due slice di interi con tale struttura
func toInt(s1 string, s2 string) (data1 []int, data2 []int) {
	d1 := strings.Split(s1, "/")
	d2 := strings.Split(s2, "/")
	for i := range d1 {
		a, _ := strconv.Atoi(d1[i])
		b, _ := strconv.Atoi(d2[i])
		data1 = append(data1, a)
		data2 = append(data2, b)
	}
	return data1, data2
}

// funzione che confronta due date in formato stringa, restituisce true se la prima data è maggiore della seconda, false altrimenti
func cmp(s1 string, s2 string) bool {
	data1, data2 := toInt(s1, s2)
	if data1[2] > data2[2] {
		return true
	} else if data1[2] == data2[2] {
		if data1[0] > data2[0] {
			return true
		} else if data1[0] == data2[0] {
			if data1[1] > data2[1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

// funzione che legge da file di input una serie di comandi e li salva
func readLines() [][]string {
	file, err := os.Open("input-1.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var comandi [][]string
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "\t")
		comandi = append(comandi, split)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
	}

	return comandi
}

// Stampa l'albero a sommario
func stampaAlberoASommario(node *node, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	fmt.Print("*")

	if node == nil {
		fmt.Println()
		return
	}

	fmt.Println(node.libro.data)

	if node.sx != nil || node.dx != nil {
		stampaAlberoASommario(node.sx, spaces+1)
		stampaAlberoASommario(node.dx, spaces+1)
	}
}

// Cerca nodo con certa data e ne stampa titolo e autore
func bist_searchITER(r *node, data string) {
	res := cmp(data, r.libro.data)
	var supp *node = r
	for supp != nil && supp.libro.data != data {
		res = cmp(data, supp.libro.data)
		if !res {
			supp = supp.sx
		} else {
			supp = supp.dx
		}
	}

	if supp == nil {
		fmt.Println("Nessun bestseller con questa data")
	} else {
		fmt.Printf("%s - %s - %s \n", supp.libro.data, supp.libro.titolo, supp.libro.autore)
	}
}

// Cerca e stampa titoli e autori dei bestseller compresi tra due date
func bist_range(r *node, data1 string, data2 string) {
	fmt.Println("{")
	res := cmp(data1, r.libro.data)
	res2 := cmp(data2, r.libro.data)
	var supp *node = r
	for supp != nil {
		res = cmp(data1, supp.libro.data)
		res2 = cmp(data2, supp.libro.data)
		if !res && res2 {
			fmt.Printf(" %s - %s - %s \n", supp.libro.data, supp.libro.titolo, supp.libro.autore)
		}

		if res || res2 {
			supp = supp.dx
		} else {
			supp = supp.sx
		}

	}
	fmt.Println("}")

}

// Elabora comandi letti effettuando le operazioni necessarie
func elaboraComandi(comandi [][]string, r **node) {
	for i, v := range comandi {
		if i == 0 {
		}
		if v[0] == "+" {
			var dati item
			dati.titolo = v[1]
			dati.autore = v[2]
			dati.editore = v[3]
			dati.data = v[4]
			dati.genere = v[5]
			n := newNode(dati)
			if i == 0 {
				*r = newNode(dati)
			} else {
				bist_insert(*r, n.libro)
			}
		} else if v[0] == "?" {
			bist_searchITER(*r, v[1])
		} else if v[0] == "-" {
			bist_range(*r, v[1], v[2])
		} else {
			return
		}
	}
}
func main() {
	t := &tree{nil}
	comandi := readLines()
	elaboraComandi(comandi, &t.root)
}
