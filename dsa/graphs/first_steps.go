package graphs

import "fmt"

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = []string{}
	}
}

func (g *Graph) AddEdge(vertex1 string, vertex2 string) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)

	g.adjacencyList[vertex1] = append(
		g.adjacencyList[vertex1],
		vertex2,
	)

	g.adjacencyList[vertex2] = append(
		g.adjacencyList[vertex2],
		vertex1,
	)
}

func (g *Graph) BFS(start string) []string {
	if _, exists := g.adjacencyList[start]; !exists {
		return []string{}
	}

	visited := make(map[string]bool)
	queue := []string{start}
	result := []string{}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current)

		for _, neighbour := range g.adjacencyList[current] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	return result
}

func (g *Graph) DFS(start string) []string {
	if _, exists := g.adjacencyList[start]; !exists {
		return []string{}
	}

	visited := make(map[string]bool)
	result := []string{}

	var visit func(string)

	visit = func(vertex string) {
		if visited[vertex] {
			return
		}

		visited[vertex] = true
		result = append(result, vertex)

		for _, neighbour := range g.adjacencyList[vertex] {
			visit(neighbour)
		}
	}

	visit(start)

	return result
}

func main() {
	graph := NewGraph()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")

	fmt.Println("BFS:", graph.BFS("A"))
	fmt.Println("DFS:", graph.DFS("A"))
}
