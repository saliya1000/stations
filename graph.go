package main

import "sort"

func BuildGraph(stations map[string]Station, connections [][2]string) map[string][]string {
	graph := make(map[string][]string)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}



func FindAllPaths(graph map[string][]string, start, end string, maxDepth int) [][]string {
	var paths [][]string
	var dfs func(path []string, visited map[string]bool)

	dfs = func(path []string, visited map[string]bool) {
		last := path[len(path)-1]
		if len(path) > maxDepth {
			return
		}
		if last == end {
			paths = append(paths, append([]string{}, path...))
			return
		}
		for _, neighbor := range graph[last] {
			if !visited[neighbor] || neighbor == end {
				visited[neighbor] = true
				dfs(append(path, neighbor), visited)
				visited[neighbor] = false
			}
		}
	}

	visited := map[string]bool{start: true}
	dfs([]string{start}, visited)

	// Sort and keep only reasonably short paths (up to 4 steps total)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// var shortPaths [][]string
	// for _, p := range paths {
	// 	if len(p) <= 4 {
	// 		shortPaths = append(shortPaths, p)
	// 	}
	// }
	return paths
	//shortPaths
}
