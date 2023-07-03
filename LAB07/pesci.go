//ARITMETICA DEI PESCI: Advent of Code nel 2021, giorno 18
//numero pesce del tipo: 	[[1,2],3] rappresentato con un albero le cui foglie sono gli effettivi numeri; un numero è composto da una coppia di due elementi, ogni elemento può essere un numero intero oppure un'altra coppia

//come passare da un numero-pesce a un albero e viceversa?

package main

import (
	"fmt"
	"strconv"
)

//definite un tipo in go per rappresentare i numeri pesce come alberi
type node struct {
	val        int
	up, sx, dx *node
}

type tree struct {
	root *node
}

//crea nuovo nodo albero
func createNode(val int) *node {
	return &node{val, nil, nil, nil}
}

// [2.1] I numeri secondo i pesci
//input: stringa che rappresenta un numero pesce; output: costruisce albero
func readTree(fishnum string) *node {
	node := createNode(-1) //sarà la radice dell'albero costruito
	support := node        //nodo di supporto

	for i := 1; i < len(fishnum); i++ {
		c := fishnum[i]          //analizzo carattere per carattere
		actual := createNode(-1) //creo nodo che sarà quello dell'albero definitivo

		//quattro casi:
		if c == '[' {
			//vuol dire creare un nuovo livello di profondità con un nodo che indirizza
			actual.up = support    //punta a quello sopra
			if support.sx == nil { //se figlio sinistro ancora vuoto riempio quello
				support.sx = actual
				support = support.sx
			} else { //altrimenti riempio figlio destro
				support.dx = actual
				support = support.dx
			}

		} else if c == ']' {
			support = support.up //risalgo di un livello, ho chiuso una coppia

		} else if c == ',' {
			continue //non mi interessa

		} else { //allora è un numero
			actual.val, _ = strconv.Atoi(string(c))
			actual.up = support
			if support.sx == nil {
				support.sx = actual
			} else {
				support.dx = actual
			}
		}

	}

	return node

}

//input: albero che rappresenta numero-pesce; output: numero-pesce con parentesi
func printTree(node *node) {
	if node == nil {
		return
	}

	if node.val != -1 {
		fmt.Print(node.val)
	} else {
		fmt.Print("[")
		printTree(node.sx)
		fmt.Print(",")
		printTree(node.dx)
		fmt.Print("]")
	}
}

//input: albero che rappresenta un numero-pesce; output: slice con le sue foglie, visitate da sinistra a destra.
func treeToSlice(node *node, num *[]int) {
	//visita in ordine anticipato con riempimento della slice
	if node == nil {
		return
	}
	if node.val != -1 {
		*num = append(*num, node.val)
	}

	treeToSlice(node.sx, num)
	treeToSlice(node.dx, num)
}

//<--------------------------------------------------------------------------------------------->
//[2.2] Somma di numeri-pesce
//La somma di due numeri-pesce x e y è data dalla coppia [x,y]. Ad esempio, [1,2] + [[3,4],5] vale [[1,2],[[3,4],5]].

//TESTING
func main() {
	node := readTree("[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]")

	num := make([]int, 0)
	printTree(node)
	fmt.Println()
	treeToSlice(node, &num)
	fmt.Println(num)
}
