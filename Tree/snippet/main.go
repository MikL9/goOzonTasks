package main

import (
	"fmt"
	"strings"
)

type Node struct {
	id       int
	text     string
	children []*Node
}

func printTree(node *Node, prefix string) {
	fmt.Println(node.text)
	for i, child := range node.children {
		newPrefix := prefix
		if i < len(node.children)-1 {
			fmt.Print(prefix + "|--")
			newPrefix += "|  "
		} else {
			fmt.Print(prefix + "`--")
			newPrefix += "   "
		}
		printTree(child, newPrefix)
	}
}

func main() {
	nodes := map[int]*Node{
		-1: {id: -1, text: "root"},
		22: {id: 22, text: "How are you?"},
		75: {id: 75, text: "I'm fine. Thank you."},
		26: {id: 26, text: "So-so"},
		45: {id: 45, text: "What's wrong?"},
		72: {id: 72, text: "Maybe I got sick"},
		81: {id: 81, text: "I wish you a speedy recovery!"},
		97: {id: 97, text: "Stick it!"},
		2:  {id: 2, text: "Thanks"},
		47: {id: 47, text: "I also got sick recently."},
		25: {id: 25, text: "Hi!"},
		82: {id: 82, text: "Bye"},
		17: {id: 17, text: "Good day!"},
		84: {id: 84, text: "Ciao!"},
		29: {id: 29, text: "Visit the doctor"},
	}

	// Build the tree
	parents := map[int]int{
		75: 22,
		84: 82,
		26: 22,
		45: 26,
		22: -1,
		72: 45,
		81: 72,
		97: 26,
		2:  97,
		47: 72,
		25: -1,
		82: -1,
		17: 82,
		29: 72,
	}
	for id, parent := range parents {
		nodes[parent].children = append(nodes[parent].children, nodes[id])
	}

	// Print the tree
	for _, child := range nodes[-1].children {
		printTree(child, "")
		fmt.Println(strings.Repeat("-", 20))
	}
}
