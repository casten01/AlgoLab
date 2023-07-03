// 1. FORESTA
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type oggetto struct {
	prezzo int
	a      string
	b      string
	op     string
}

// Legga da stdin gli indizi uno per riga nel formato oggetto:indizio, crei un oggetto per ogni indizio e li memorizzi in una mappa

func createObj() *oggetto {
	return &oggetto{0, "", "", ""}
}

func leggiInput() map[string]*oggetto {
	listino := make(map[string]*oggetto)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := scanner.Text()
		indizio := strings.Split(riga, ":")
		var operandi []string = make([]string, 0)
		if len(indizio[1]) > 1 {
			operandi = make([]string, 3)
			operandi = strings.Split(indizio[1], " ")
		} else {
			operandi = append(operandi, indizio[1])
		}
		obj := createObj()
		if len(operandi) > 1 {
			obj.a = operandi[0]
			obj.op = operandi[1]
			obj.b = operandi[2]
		} else {
			obj.prezzo, _ = strconv.Atoi(operandi[0])
		}
		listino[indizio[0]] = obj

	}

	return listino
}

func main() {
	listino := leggiInput()
	for k, v := range listino {
		if (*v).prezzo != 0 {
			fmt.Printf("%s: %d \n", k, (*v).prezzo)
			continue
		}
		fmt.Printf("%s: %s %s %s \n", k, (*v).a, (*v).op, (*v).b)

	}
}
