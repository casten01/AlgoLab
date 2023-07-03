//SOMMA della COPPIA: si descriva un algoritmo che dato un intero k e un albero di ricerca binario (nodi con chiavi intere), stabilisce se esistono due nodi nell'albero le cui chiavi hanno somma k

package main

import "fmt"

type treeNode struct {
	left  *treeNode
	right *treeNode
	item  int
}

type tree struct {
	root *treeNode
}

//FUNZIONI AUSILIARIE
func cmp(a int, b int) int {
	return a - b
}

/*Crea un nodo dell'albero*/
func newNode(val int) *treeNode {
	return &treeNode{nil, nil, val}
}

//RICERCA DI UNA CHIAVE:
/* Versione ricorsiva*/
func bist_searchRIC(p *treeNode, item int) string {
	if p == nil {
		return "Elemento non trovato"
	}
	res := cmp(item, p.item)

	if res == 0 {
		return "Elemento trovato"
	} else if res < 0 {
		return bist_searchRIC(p.left, item)
	} else {
		return bist_searchRIC(p.right, item)
	}
}

/*Versione iterativa*/
func bist_searchITER(p *treeNode, item int) string {
	var res int = cmp(item, p.item)
	var supp *treeNode = p
	for supp != nil && res != 0 {
		res = cmp(item, supp.item)
		if res < 0 {
			supp = supp.left
		} else if res > 0 {
			supp = supp.right
		}
	}

	if supp == nil {
		return "Elemento non trovato"
	} else {
		return "Elemento trovato"
	}
}

//IDENTIFICARE IL PADRE DI UN NODO
func bist_searchparent(r *treeNode, item int, parent *treeNode, p *treeNode) string {
	*p = *r
	p.item = 100
	var res int = cmp(item, p.item)
	if r == nil {
		return "L'albero non è istanziato"
	}
	for res != 0 {
		res = cmp(item, p.item)
		if p.left == nil && p.right == nil && res != 0 {
			*parent = *p
			p = nil
			break
		}
		if res < 0 {
			*parent = *p
			*p = *p.left
		} else if res > 0 {
			*parent = *p
			*p = *p.right
		}
	}

	if p == nil {
		return "Elemento non trovato"
	} else {
		return "Elemento trovato"
	}
}

func bist_searchparent2(r **treeNode, item int) *treeNode {
	p := *r
	var parent *treeNode = nil

	for p != nil {
		res := cmp(item, p.item)
		if res == 0 {
			return p
		} else if res < 0 {
			parent = p
			p = p.left
		} else {
			parent = p
			p = p.right
		}
	}

	return parent
}

//INSERIRE UN NUOVO ELEMENTO
func bist_insert(r *treeNode, item int) string {
	if r == nil {
		return "L'albero è vuoto"
	}
	res := bist_searchparent2(&r, item)
	if res.item == item {
		return "Elemento già presente"
	}

	if cmp(item, res.item) < 0 {
		res.left = newNode(item)
	} else {
		res.right = newNode(item)
	}
	return "Elemento inserito correttamente"
}

//STAMPA L'ALBERO
func stampaAlberoASommario(node *treeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	fmt.Print("*")

	if node == nil {
		fmt.Println()
		return
	}

	fmt.Println(node.item)

	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}

}

/*
//CANCELLA ELEMENTO CON UNA CERTA CHIAVE
func bist_delete(r *treeNode, item int) {
	if *r == nil {
		return
	}

	// Trova il nodo da eliminare
	var parent *treeNode = nil
	p := *r
	for p != nil && p.item != item {
		parent = p
		if item < p.item {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil {
		// L'elemento non è presente nell'albero
		return
	}

	// Caso 1: Il nodo da eliminare è una foglia
	if p.left == nil && p.right == nil {
		if parent == nil {
			*r = nil
		} else if parent.left == p {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if p.left != nil && p.right != nil {
		// Caso 3: Il nodo da eliminare ha due figli
		successor := findMinNode(p.right)
		p.item = successor.item
		bist_delete(&p.right, successor.item)
	} else {
		// Caso 2: Il nodo da eliminare ha un solo figlio
		var child *treeNode
		if p.left != nil {
			child = p.left
		} else {
			child = p.right
		}

		if parent == nil {
			*r = child
		} else if parent.left == p {
			parent.left = child
		} else {
			parent.right = child
		}
	}
}

// Trova il nodo con il valore minimo nell'albero
func findMinNode(p *treeNode) *treeNode {
	if p == nil {
		return nil
	}

	for p.left != nil {
		p = p.left
	}

	return p
}
*/

func main() {
	t := &tree{nil}
	t.root = newNode(10)
	t.root.left = newNode(6)
	t.root.right = newNode(28)

	t.root.left.left = newNode(3)

	t.root.left.right = newNode(8)
	t.root.left.right.left = newNode(7)
	t.root.left.right.right = newNode(9)

	t.root.right.left = newNode(16)
	t.root.right.right = newNode(30)

	// p := newNode(0)
	//parent := newNode(0)

	// result := bist_searchparent(t.root, 9, parent, p)
	// if result != "Elemento trovato" {
	// 	fmt.Println("Attaccare l'eventuale nodo al nodo con item: ", parent.item)
	// } else {
	// 	fmt.Printf("%s: il padre di %d ha item %d \n", result, p.item, parent.item)
	// }

	fmt.Println(bist_insert(t.root, 16))
	fmt.Println(bist_insert(t.root, 42))

	stampaAlberoASommario(t.root, 0)
	//fmt.Println(bist_searchparent(t.root, 30, parent, p))

	//fmt.Println(bist_searchITER(t.root, 30))
	//fmt.Println(bist_searchRIC(t.root, 12))

}
