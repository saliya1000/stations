package main

import "fmt"

func FilterDisjointPaths(paths [][]string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	pathsShareIntermediates := func(path1, path2 []string) bool {
		for i := 1; i < len(path1)-1; i++ {
			for j := 1; j < len(path2)-1; j++ {
				if path1[i] == path2[j] {
					return true
				}
			}
		}
		return false
	}

	bestPaths := [][]string{paths[0]}
	bestTurns := 1 << 30

	var tryPaths func([][]string, int, [][]string)
	tryPaths = func(remaining [][]string, index int, current [][]string) {
		if len(current) > 0 {
			turns := CalculateMinTurns(current, 10)
			if turns < bestTurns {
				bestTurns = turns
				bestPaths = make([][]string, len(current))
				copy(bestPaths, current)
			}
		}

		for i := index; i < len(remaining); i++ {
			isDisjoint := true
			for _, selectedPath := range current {
				if pathsShareIntermediates(remaining[i], selectedPath) {
					isDisjoint = false
					break
				}
			}
			if isDisjoint {
				newCurrent := append(current, remaining[i])
				tryPaths(remaining, i+1, newCurrent)
			}
		}
	}

	tryPaths(paths, 0, [][]string{})

	fmt.Println("Best paths found:")
	for i, path := range bestPaths {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
	fmt.Println("-------------------------------")
	return bestPaths
}
