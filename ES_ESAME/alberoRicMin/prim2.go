//Coda con priorità -> slice con indice del minimo

package main

import "math"

type Coda struct {
	lista []int
	min   int
}

type Arco struct {
	v1, v2 int
	peso   int
}
type Graph struct {
	archi []Arco
	n     int
}

const MAX = math.MaxInt32

//OPERAZIONI PER LA GESTIONE DELLA CODA
/*Inizializza coda della grandezza del numero di vertici con valori ~ infinito*/
func riempiCoda(coda *Coda, n_vertici int) {
	coda.lista = make([]int, n_vertici)

	for i := range coda.lista {
		coda.lista[i] = MAX
	}
}

/*Cambia valore della posizione idx della coda e cambia eventualmente il minimo*/
func changeCoda(coda *Coda, idx int, val int) {
	coda.lista[idx] = val
	if val < coda.lista[coda.min] {
		coda.min = idx
	}
}

/*Elimina il minimo dalla coda e ne restituisce l'indice, cambia con valore negativo che indica che l'elemento non è già stato preso */
func eliminaMin(coda *Coda) (idx int) {
	idx = coda.min
	coda.lista[coda.min] = -1
	for i, v := range coda.lista {
		if v < coda.lista[coda.min] && v > 0 {
			coda.min = i
		}
	}
	return idx
}
/*Controlla che la coda non sia vuota, tutti gli elementi corrispondono a -1*/
func isEmpty(coda *Coda) bool {
	for _, v := range coda.lista {
		if v != -1 {
			return false
		}
	}
	return true
}

//CREA UN ARCO
func creaArco(v1 int, v2 int, peso int) (e Arco) [
	e.v1 = v1
	e.v2 = v2
	e.peso = peso

	return e
]

func prim(g Graph) (tree []Arco) {
	var c *Coda
	vicino := make([]int, g.n)
	d := make([]int, g.n)

	riempiCoda(c, g.n)
	for i := range d {
		d[i] = MAX
	}

	for isEmpty(c) {
		y := eliminaMin(c)

		if d[y] != MAX {
			x := vicino[y]
			e := creaArco(x, y, d[y])
			tree = append(tree, e)
		}

		for _, arco := range g.archi {
			if arco.v1 == y || arco.v2 == y {
				//basta
			}
		}
	}

	return tree
}

func main() {

}
