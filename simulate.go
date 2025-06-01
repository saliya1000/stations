package main

import (
	"fmt"
	"sort"
	"strings"
)

func CalculateMinTurns(chosenPaths [][]string, numTrains int) int {
	if len(chosenPaths) == 0 {
		return 0
	}

	maxL := 0
	for _, path := range chosenPaths {
		if len(path) > 1 {
			L := len(path) - 1
			if L > maxL {
				maxL = L
			}
		}
	}
	if maxL == 0 {
		return 1
	}

	T := maxL + 1
	for {
		total := 0
		for _, path := range chosenPaths {
			if len(path) > 1 {
				L := len(path) - 1
				if T > L {
					total += (T - L)
				}
			}
		}
		if total >= numTrains {
			return T
		}
		T++
	}
}

func CreateTrainAssignments(chosenPaths [][]string, numTrains, bestT int) []PathAssignment {
	var assignments []PathAssignment
	for i, path := range chosenPaths {
		if len(path) > 1 {
			L := len(path) - 1
			for t := 1; t <= bestT-L; t++ {
				assignments = append(assignments, PathAssignment{PathIndex: i, EntryTurn: t})
			}
		}
	}
	sort.Slice(assignments, func(i, j int) bool {
		if assignments[i].EntryTurn == assignments[j].EntryTurn {
			return assignments[i].PathIndex < assignments[j].PathIndex
		}
		return assignments[i].EntryTurn < assignments[j].EntryTurn
	})

	if len(assignments) > numTrains {
		assignments = assignments[:numTrains]
	}
	// fmt.Println("Train assignments:")
	// for i, path := range assignments {
	// 	fmt.Printf("Train %d: %v\n", i+1, path)
	// 	//fmt.Println(len(path))
	// }

	// //fmt.Println(len(assignments))
	// fmt.Println("-------------------------------")
	return assignments
}

func SimulateTrainMovements(trains []*Train, totalTurns int) []string {
	var output []string
	for turn := 1; turn <= totalTurns; turn++ {
		var moves []string
		for _, train := range trains {
			if turn < train.EntryTurn || turn >= train.EntryTurn+len(train.Path) {
				continue
			}
			pos := turn - train.EntryTurn + 1
			if pos < len(train.Path) {
				moves = append(moves, fmt.Sprintf("%s-%s", train.ID, train.Path[pos]))
			}
		}
		if len(moves) > 0 {
			sort.Strings(moves)
			output = append(output, strings.Join(moves, " "))
		}
	}

	return output
}

func ScheduleTrains(network Network, start, end string, numTrains int) ([]string, error) {
	paths, err := FindPaths(&network, start, end)
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, NewValidationError("error: no valid path", start, end)
	}

	disjointPaths := FilterDisjointPaths(paths)

	if len(disjointPaths) == 0 {
		return nil, NewValidationError("error: no disjoint paths", start, end)
	}

	minTurns := CalculateMinTurns(disjointPaths, numTrains)
	assignments := CreateTrainAssignments(disjointPaths, numTrains, minTurns)
	uniquePathIndices := make(map[int]bool)
	for _, assignment := range assignments {
		uniquePathIndices[assignment.PathIndex] = true
	}
	fmt.Printf("Number of distinct paths used by trains: %d\n", len(uniquePathIndices))
	fmt.Println("-------------------------------")

	if len(uniquePathIndices) > 0 {
		fmt.Println("Details of used paths:")
		sortedPathIndices := make([]int, 0, len(uniquePathIndices))
		for index := range uniquePathIndices {
			sortedPathIndices = append(sortedPathIndices, index)
		}
		sort.Ints(sortedPathIndices) // Sort the indices for consistent output

		for _, pathIndex := range sortedPathIndices {
			// It's good practice to check if pathIndex is within bounds of disjointPaths
			if pathIndex >= 0 && pathIndex < len(disjointPaths) {
				fmt.Printf("  - Path %d: %v\n", pathIndex+1, disjointPaths[pathIndex])
				// fmt.Println("-------------------------------")
			} else {
				// This case should ideally not happen if logic is correct, but good for robustness
				fmt.Printf("error: invalid path index %d encountered\n", pathIndex)
				fmt.Println("-------------------------------")
			}
		}
		fmt.Println("-------------------------------")
	}

	trains := make([]*Train, len(assignments))
	for i, assignment := range assignments {
		trains[i] = &Train{
			ID:        fmt.Sprintf("T%d", i+1),
			Path:      disjointPaths[assignment.PathIndex],
			EntryTurn: assignment.EntryTurn,
		}
	}

	return SimulateTrainMovements(trains, minTurns), nil
}
