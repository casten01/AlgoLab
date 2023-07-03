//ALBERI BINARI - implementazione con puntatori ai figli destro e sinistro

package main

import (
	"fmt"
	"math/rand"
)

const N = 10 //numero di nodi dell'albero

//nodi albero binario
type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

//genera a caso una sequenza di interi di lunghezza fissata con la macro N e la memorizzi in una slice
func randomSeq() []int {
	seq := make([]int, N)

	for i := 0; i < N; i++ {
		var n = rand.Intn(100)
		seq[i] = n
	}

	return seq
}

//crea un nuovo nodo
func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}

//costruisce un albero a partire dalla slice
func createTree(seq []int) bitree {
	t := &bitree{nil}

	t.root = &bitreeNode{nil, nil, seq[0]}

	t.root.left = newNode(seq[1])
	t.root.right = newNode(seq[2])
	t.root.left.left = newNode(seq[3])
	t.root.left.left.left = newNode(seq[7])
	t.root.left.left.right = newNode(seq[4])
	t.root.right.left = newNode(seq[5])
	t.root.right.right = newNode(seq[6])
	t.root.right.left.left = newNode(seq[8])
	t.root.right.right.left = newNode(seq[9])

	return *t
}

//ORDINE SIMMETRICO
func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(" ", node.val)
	inorder(node.right)
}

//ORDINE ANTICIPATO
func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(" ", node.val)
	preorder(node.left)
	preorder(node.right)
}

//ORDINE POSTICIPATO
func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(" ", node.val)
}

//STAMPA DI ALBERI A SOMMARIO: figlio mancante => riga vuota
// utilizzeremo la visita in ordine anticipato, il numero di spazi può essere dato come parametro della funzione
// andrà chiamata -> stampaAlberoASommario ( root , 0 );
func stampaAlberoASommario(node *bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	fmt.Print("*")

	if node == nil {
		fmt.Println()
		return
	}

	fmt.Println(node.val)

	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}

}

//UNA STAMPA ALTERNATIVA: albero stampato su una sola riga,
func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}

	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

//DAL VETTORE ALL'ALBERO

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}

	nodo := newNode(a[i])

	root = nodo
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}

//MAIN
func main() {
	//CREA ALBERO
	seq := randomSeq()
	t := createTree(seq)

	//VISITE
	fmt.Print("Ordine simmetrico: ")
	inorder(t.root)
	fmt.Println()
	fmt.Print("Ordine posticipato: ")
	postorder(t.root)
	fmt.Println()
	fmt.Print("Ordine anticipato: ")
	preorder(t.root)

	fmt.Println()

	//STAMPA A SOMMARIO
	stampaAlberoASommario(t.root, 0)

	//STAMPA ALTERNATIVA
	stampaAlbero(t.root)

	fmt.Println()
	//DAL VETTORE ALL'ALBERO
	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	root2 := arr2tree(a, 0)
	stampaAlberoASommario(root2, 0)
}
