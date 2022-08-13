package main

import "fmt"

type grid struct {
	matrix [][]string
}

func (g *grid) bfsSearch(startRow int, startCol int, endRow int, endCol int) {
	rowQueue := []int{ startRow }
	colQueue := []int{ startCol }

	var parentRowIds []int
	var parentColIds []int


	for i := 0; i < len(g.matrix); i++ {
		parentRowIds = append(parentRowIds, -1)
	}

	for i := 0; i < len(g.matrix[0]); i++ {
		parentColIds = append(parentColIds, -1)
	}

	var visitedRows []int
	var visitedColumns []int

	for len(rowQueue) > 0 {
		currentRow := rowQueue[0]
		currentColumn := colQueue[0]

		rowQueue = rowQueue[1:]
		colQueue = colQueue[1:]

		visitedRows = append(visitedRows, currentRow)
		visitedColumns = append(visitedColumns, currentColumn)


		//for _, neighbour := range nodeNeighbours {
		//	if !visited[neighbour.name] {
		//		queue = append(queue, neighbour)
		//		visited[neighbour.name] = true
		//		parentNodeIds[neighbour.id] = currentNode.id
		//	}
		//}
	}

}

//func (g *graph) gridShortestPath(startNode node, endNode node) []int {
//	parentPathIds := g.bfsSearch(startNode)
//	var path []int
//
//	fmt.Println(parentPathIds)
//
//	for currentNodeId := endNode.id; currentNodeId != -1; currentNodeId = parentPathIds[currentNodeId] {
//		fmt.Println(currentNodeId)
//		path = append(path, currentNodeId)
//	}
//
//	var pathReversed []int
//
//	for i := len(path) - 1; i >= 0; i-- {
//		pathReversed = append(pathReversed, path[i])
//	}
//
//	if pathReversed[0] != startNode.id {
//		return []int{}
//	}
//
//	return pathReversed
//}

func main()  {
	dungeonGrid := [][]string{
		{ "S", ".", ".", "#", ".", ".", "." },
		{ ".", "#", ".", ".", ".", "#", "." },
		{ ".", "#", ".", ".", ".", ".", "." },
		{ ".", ".", "#", "#", ".", ".", "." },
		{ "#", ".", "#", "E", ".", "#", "." },
	}

	g := grid{ matrix: dungeonGrid }
	fmt.Println(g)
}


