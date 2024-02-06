package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ID       int
	ParentId int
	Content  string
	Children []*Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		allNodes := map[int]*Node{}
		links := map[int]int{}

		for j := 1; j <= n; j++ {
			scanner.Scan()
			row := scanner.Text()
			rowParts := strings.SplitN(row, " ", 3)
			var node Node
			id, _ := strconv.Atoi(rowParts[0])
			node.ID = id
			parent, _ := strconv.Atoi(rowParts[1])
			node.ParentId = parent
			node.Content = rowParts[2]
			allNodes[node.ID] = &node
			links[node.ID] = node.ParentId
		}
		var node Node
		node.ID = -1
		node.Content = "DA KTO YA?"
		allNodes[node.ID] = &node
		for nodeId, parent := range links {
			allNodes[parent].Children = append(allNodes[parent].Children, allNodes[nodeId])
		}

		for _, child := range allNodes[-1].Children {
			makeLinks(child, "")
		}
	}
}

func makeLinks(node *Node, palka string) {
	fmt.Println(node.Content)
	for level, child := range node.Children {
		palki := palka
		fmt.Print(palki + "|--")
		if level < len(node.Children)-1 {
			palki += "|  "
		} else {
			palki += "   "
		}
		makeLinks(child, palki)
	}
}
