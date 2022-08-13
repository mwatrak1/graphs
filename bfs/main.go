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

func (g *graph) bfsSearch(startNode node) {
	queue := []node{ startNode }

	visited := make(map[string]bool)
	visited[startNode.name] = true

	queue = append(queue, startNode)

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		nodeNeighbours := g.adjacencyList[currentNode]

		for _, neighbour := range nodeNeighbours {
			if !visited[neighbour.name] {
				fmt.Println(neighbour.name)
				queue = append(queue, neighbour)
				visited[neighbour.name] = true
			}
		}
	}

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
	g.bfsSearch(a)
}
