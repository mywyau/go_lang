package graphs

func journeyToMoon(n int32, astronaut [][]int32) int64 {
	// Build an adjacency list.
	graph := make([][]int32, n)

	for _, pair := range astronaut {
		a := pair[0]
		b := pair[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)

	var result int64
	var astronautsSeen int64

	for astronautID := int32(0); astronautID < n; astronautID++ {
		if visited[astronautID] {
			continue
		}

		countrySize := dfsCountrySize(
			astronautID,
			graph,
			visited,
		)

		// Everyone in this country can pair with everyone
		// from previously discovered countries.
		result += countrySize * astronautsSeen

		astronautsSeen += countrySize
	}

	return result
}

func dfsCountrySize(
	astronaut int32,
	graph [][]int32,
	visited []bool,
) int64 {
	visited[astronaut] = true

	var size int64 = 1

	for _, neighbour := range graph[astronaut] {
		if !visited[neighbour] {
			size += dfsCountrySize(neighbour, graph, visited)
		}
	}

	return size
}