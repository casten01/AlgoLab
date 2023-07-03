package main

import (
    "fmt"
    "strings"
    "strconv"
)

type listNode struct {
  item int
  next *listNode
}

type pila struct {
    top *listNode
}

func pop(p *pila) int  {
  var x *listNode = p.top
  p.top = x.next
  return x.item
}

func push(p *pila, x int) {
  var node *listNode = new(listNode)
  node.item = x
  node.next = p.top
  p.top = node
}

func main() {
  //test funzioni pila?
  var p pila = pila{nil}
  fmt.Println(valuta("5 3 - 2 *", &p))
}

func valuta(espressione string, p *pila) int  {
  var res int
  var elementi []string = strings.Split(espressione, " ")

  ...

  push(p, n)

  ...

  b := pop(p)
  a := pop(p)

      
}


/* Note:

In questo programma, la pila è:

1. definita come struttura di dati astratta:
  - è definito un tipo "pila" (riga 14)
  - sono definite delle funzioni generali "pop", "push", ...

2. usata come struttura di dati astratta:
  - nel main si dichiara una variabile p di tipo pila
  - la variabile p viene modificata solo mediante l'uso delle funzioni "pop", "push"

3. implementata mediante una lista concatenata:
  - la funzione push effettua un inserimento in testa
  - la funzione pop effettua una rimozione in testa


Possibili miglioramenti:
- mancano altre funzioni tipiche della pila (isEmpty, top)
- se riceve una pila vuota, la funzione "pop" dà un panic error alla riga 26, 
poiché si cerca di accedere al campo "next" della struttura "p.top" che invece ha valore "nil".

In questo programma non si verifica mai questa situazione, ma implementando la struttura di dati astratta,
sarebbe opportuno prendere in considerazione anche questo caso.
(si vedano le note al programma tag-pilaADT.go )

*/