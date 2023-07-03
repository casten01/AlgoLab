/* CONTO CORRENTE: Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un
conto corrente. Il programma legge poi da standard input una serie di numeri interi che rappresentano
spese da addebitare sul conto e stampa il saldo finale. La lettura si interrompe quando il
saldo è <=0. */

package main

import "fmt"

func main() {
	var saldo int
	var spesa int
	fmt.Print("Inserire saldo iniziale: ")
	fmt.Scan(&saldo)

	for saldo >= 0 {
		fmt.Print("spesa addebitata da: ")
		fmt.Scan(&spesa)
		saldo -= spesa
		fmt.Println("saldo attuale: ", saldo)
	}
}
