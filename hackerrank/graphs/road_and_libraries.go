package graphs

func roadsAndLibraries(
	n int32,
	cLib int32,
	cRoad int32,
	cities [][]int32,
) int64 {

	// If a library is no more expensive than a road,
	// build one library in every city.
	if cLib <= cRoad {
		return int64(n) * int64(cLib)
	}

	// Cities are numbered from 1 to n, so index 0 is unused.
	graph := make([][]int32, n+1)

	// The roads are bidirectional, so add both directions.
	for _, road := range cities {
		cityA := road[0]
		cityB := road[1]

		graph[cityA] = append(graph[cityA], cityB)
		graph[cityB] = append(graph[cityB], cityA)
	}

	visited := make([]bool, n+1)
	var components int64

	// Check every city, including isolated cities.
	for city := int32(1); city <= n; city++ {
		if visited[city] {
			continue
		}

		// This unvisited city begins a new component.
		components++

		// Iterative depth-first search.
		stack := []int32{city}
		visited[city] = true

		for len(stack) > 0 {
			lastIndex := len(stack) - 1
			current := stack[lastIndex]
			stack = stack[:lastIndex]

			for _, neighbour := range graph[current] {
				if !visited[neighbour] {
					visited[neighbour] = true
					stack = append(stack, neighbour)
				}
			}
		}
	}

	libraryCost := components * int64(cLib)
	numberOfRoads := int64(n) - components
	roadCost := numberOfRoads * int64(cRoad)

	return libraryCost + roadCost
}
