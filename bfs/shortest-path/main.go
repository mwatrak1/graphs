package main

import "fmt"

type node struct {
	id int
	name string
}

func newNode(id int, name string) node  {
	n := node{ id: id, name: name }
	return n
}

type graph struct {
	adjacencyList map[node][]node
}

func (g *graph) getNode(nodeId int) {

}

func (g *graph) bfsSearch(startNode node) []int  {
	queue := []node{ startNode }

	var parentNodeIds []int
	nodeCount := len(g.adjacencyList)

	for i := 0; i < nodeCount; i++ {
		parentNodeIds = append(parentNodeIds, -1)
	}

	visited := make(map[string]bool)
	visited[startNode.name] = true

	queue = append(queue, startNode)

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		nodeNeighbours := g.adjacencyList[currentNode]

		for _, neighbour := range nodeNeighbours {
			if !visited[neighbour.name] {
				queue = append(queue, neighbour)
				visited[neighbour.name] = true
				parentNodeIds[neighbour.id] = currentNode.id
			}
		}
	}

	return parentNodeIds
}

func (g *graph) shortestPath(startNode node, endNode node) []int {
	parentPathIds := g.bfsSearch(startNode)
	var path []int

	fmt.Println(parentPathIds)

	for currentNodeId := endNode.id; currentNodeId != -1; currentNodeId = parentPathIds[currentNodeId] {
		fmt.Println(currentNodeId)
		path = append(path, currentNodeId)
	}

	var pathReversed []int

	for i := len(path) - 1; i >= 0; i-- {
		pathReversed = append(pathReversed, path[i])
	}

	if pathReversed[0] != startNode.id {
		return []int{}
	}

	return pathReversed
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
	adjacencyList[b] = []node{ d }
	adjacencyList[c] = []node{ e }
	adjacencyList[d] = []node{ f }
	adjacencyList[e] = []node{  }
	adjacencyList[f] = []node{  }

	g := graph{adjacencyList: adjacencyList}
	shortestPathIds := g.shortestPath(c, e)
	fmt.Println(shortestPathIds)
}

