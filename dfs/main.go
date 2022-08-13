package main

import "fmt"

type node struct {
	name string
}

func newNode(name string) node  {
	n := node{name: name}
	return n
}

type graph struct {
	adjacencyList map[string][]node
}

func (g *graph) dfsSearch(startNodeName string) {
	adjacencyList := g.adjacencyList
	visitedNodes := make(map[string]bool)

	var dfs func(nodeName string)

	dfs = func (nodeName string) {
		if visitedNodes[nodeName] {
			return
		}

		visitedNodes[nodeName] = true

		fmt.Println(nodeName)

		for _, neighbour := range adjacencyList[nodeName] {
			dfs(neighbour.name)
		}
	}

	dfs(startNodeName)
}

func main()  {
	adjacencyList := make(map[string][]node)

	adjacencyList["A"] = []node{ newNode("B"), newNode("C") }
	adjacencyList["B"] = []node{ newNode("C") }
	adjacencyList["C"] = []node{ newNode("B"), newNode("A") }
	adjacencyList["D"] = []node{}

	g := graph{adjacencyList: adjacencyList}

	g.dfsSearch("C")
}
