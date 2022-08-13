package main

import "fmt"

type node struct {
	id int
	name string
	groupId int
}

func newNode(id int, name string) node  {
	n := node{ id: id, name: name, groupId: -1 }
	return n
}

type graph struct {
	adjacencyList map[node][]node
}

func (n *node) setGroupId(groupId int)  {
	n.groupId = groupId
}

func (g *graph) dfsSearch(startNode node) []node {
	adjacencyList := g.adjacencyList
	visitedNodes := make(map[string]bool)
	var foundNodes []node

	var dfs func(currentNode node)

	dfs = func (currentNode node) {
		if visitedNodes[currentNode.name] {
			return
		}

		visitedNodes[currentNode.name] = true

		foundNodes = append(foundNodes, currentNode)

		for _, neighbour := range adjacencyList[currentNode] {
			dfs(neighbour)
		}
	}

	dfs(startNode)

	return foundNodes
}

func (g* graph) findConnectedComponents() [][]node  {
	i := 0

	visitedNodes := make(map[string]bool)
	var componentGroups [][]node

	for currentNode, _ := range g.adjacencyList {

		componentGroup := g.dfsSearch(currentNode)
		var indexedComponentGroup []node

		for _, component := range componentGroup {

			if !visitedNodes[component.name] {
				component.setGroupId(i)
				indexedComponentGroup = append(indexedComponentGroup, component)
				visitedNodes[component.name] = true
			}

		}

		if len(indexedComponentGroup) > 0 {
			componentGroups = append(componentGroups, indexedComponentGroup)
			i++
		}

	}

	return componentGroups
}

func main()  {
	adjacencyList := make(map[node][]node)

	a := newNode(0, "A")
	b := newNode(1, "B")
	c := newNode(2, "C")
	d := newNode(3, "D")
	e := newNode(4, "E")
	f := newNode(5, "F")

	adjacencyList[a] = []node{ b, c }
	adjacencyList[b] = []node{ c }
	adjacencyList[c] = []node{ a, b }
	adjacencyList[d] = []node{ e }
	adjacencyList[e] = []node{ d }
	adjacencyList[f] = []node{}

	g := graph{adjacencyList: adjacencyList}
	connectedComponentGroups := g.findConnectedComponents()
	fmt.Println(connectedComponentGroups)
}
