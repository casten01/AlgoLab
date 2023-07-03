package main

import "fmt"

type node struct {
	val        int
	up, sx, dx *node
}

func createNode(val int) *node {
	return &node{val, nil, nil, nil}
}

func main() {
	node := createNode(-1)
	support := node
	fmt.Println(support.val)
	support.val = 10
	support.sx = createNode(10)
	support = support.sx
	support.val = 4
	fmt.Println(support.val)
	fmt.Println(node.val)

}
