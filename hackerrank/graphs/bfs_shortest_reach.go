package graphs

func bfs(
	n int32,
	m int32,
	edges [][]int32,
	start int32,
) []int32 {
	// HackerRank node numbers start at 1.
	// We therefore create n+1 positions and ignore index 0.
	graph := make([][]int32, n+1)

	// Build the undirected adjacency list.
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	// Initialise every node as unreachable.
	distances := make([]int32, n+1)

	for node := int32(1); node <= n; node++ {
		distances[node] = -1
	}

	// The distance from the start node to itself is 0.
	distances[start] = 0

	// Start BFS from the given node.
	queue := []int32{start}

	for len(queue) > 0 {
		// Take the first node from the queue.
		current := queue[0]
		queue = queue[1:]

		// Explore all neighbouring nodes.
		for _, neighbour := range graph[current] {
			// A value other than -1 means that this node
			// has already been visited.
			if distances[neighbour] != -1 {
				continue
			}

			distances[neighbour] = distances[current] + 6
			queue = append(queue, neighbour)
		}
	}

	// HackerRank does not want the starting node
	// included in the result.
	result := make([]int32, 0, n-1)

	for node := int32(1); node <= n; node++ {
		if node == start {
			continue
		}

		result = append(result, distances[node])
	}

	return result
}