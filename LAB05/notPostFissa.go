//NOTAZIONE POSTFISSA:  5 3 - 2 * equivale a (5-3)*2

//1.1 VALUTAZIONE di un'ESPRESSIONE in NOTAZIONE POSTFISSA stack -> LIFO
/*
leggi un token (operatore o numero );
se il token è un numero
	inseriscilo nella pila;
se il token è un operatore
	estrai due valori dalla pila;
applica ad essi l’operatore;
inserisci il risultato nella pila;
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//push
func push(pila []string, number string) []string {
	pila = append(pila, number)
	return pila
}

//pop
func popDoppio(pila []string) ([]string, int, int) {
	a, _ := strconv.Atoi(pila[len(pila)-2])
	b, _ := strconv.Atoi(pila[len(pila)-1])
	pila = pila[:len(pila)-2]
	return pila, a, b
}

//operazioni
func operazioni(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	fmt.Println("error")
	return 0
}

//riceve espressione in funzione post fissa e restituisce valore
func valuta(espressione string) int {
	var pila []string
	var res int
	tokens := strings.Split(espressione, " ")

	for i := 0; i < len(tokens); i++ {
		var token string = tokens[i]

		if unicode.IsDigit(rune(token[0])) {
			pila = push(pila, tokens[i])

		} else {
			var a, b int
			pila, a, b = popDoppio(pila)
			op := tokens[i]
			res = operazioni(a, b, rune(op[0]))
			pila = push(pila, strconv.Itoa(res))
		}
	}

	return res
}

//1.2 CONVERSIONE da NOTAZIONE INFISSA a NOTAZIONE POSTFISSA
/*
leggi un token;
se il token è un numero
	stampa il token;
se il token è un operatore
	inserisci il token in cima alla pila;
se il token è una parentesi aperta
	ignora il token
se il token è una parentesi chiusa
	estrai l’operatore in cima alla pila;
	stampalo;
*/

//riceve una espressione in notazione infissa e restituisce l’espressione equivalente in notazione postfissa
//( ( 5 - 3 ) * 2 ) -> 5 3 - 2 *

//pop singolo
func pop(pila []string) []string {
	op := pila[len(pila)-1]
	pila = pila[:len(pila)-1]

	fmt.Printf("%s ", op)
	return pila
}

func converti(espressione string) {
	var pila []string
	tokens := strings.Split(espressione, " ")

	for i := 0; i < len(tokens); i++ {
		var token string = tokens[i]

		if unicode.IsDigit(rune(token[0])) {
			fmt.Printf("%s ", tokens[i])
		} else if tokens[i] == ")" {
			pila = pop(pila)
		} else if tokens[i] == "(" {
			continue
		} else {
			pila = push(pila, tokens[i])
		}
	}
}

func main() {
	res := valuta("5 3 - 2 *")
	fmt.Println(res)
	converti("( ( 5 - 3 ) * 2 )")
}
