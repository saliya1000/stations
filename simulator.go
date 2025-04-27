// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func ScheduleTrains(paths [][]string, trainCount int) {
// 	type Train struct {
// 		name     string
// 		path     []string
// 		delay    int
// 		position int // index in path
// 	}

// 	var trains []Train

// 	// Sort paths by shortest first
// 	sort.Slice(paths, func(i, j int) bool {
// 		return len(paths[i]) < len(paths[j])
// 	})

// 	// Assign trains with staggered delays to avoid crowding
// 	// Only pick the first 2 paths (shortest ones)
// 	usablePaths := paths
// 	if len(paths) > 2 {
// 		usablePaths = paths[:2]
// 	}

// 	for i := 0; i < trainCount; i++ {
// 		path := usablePaths[i%len(usablePaths)]
// 		delay := i / len(usablePaths) // delay based on cycle of 2 paths
// 		trains = append(trains, Train{
// 			name:  fmt.Sprintf("T%d", i+1),
// 			path:  path[1:], // exclude start station
// 			delay: delay,
// 		})
// 	}

// 	// Debug train assignments
// 	fmt.Println("Train assignments:")
// 	for _, train := range trains {
// 		fmt.Printf("%s: %v (delay: %d)\n", train.name, train.path, train.delay)
// 	}
// 	fmt.Println("--------------------------------")

// 	finished := false
// 	turn := 0

// 	for !finished {
// 		turn++
// 		moves := []string{}
// 		occupied := make(map[string]bool)
// 		stillMoving := false

// 		for i := range trains {
// 			train := &trains[i]
// 			if train.delay > 0 {
// 				train.delay--
// 				continue
// 			}
// 			if train.position >= len(train.path) {
// 				continue
// 			}
// 			next := train.path[train.position]
// 			if !occupied[next] {
// 				moves = append(moves, fmt.Sprintf("%s-%s", train.name, next))
// 				occupied[next] = true
// 				train.position++
// 				stillMoving = true
// 			}
// 		}

// 		if len(moves) > 0 {
// 			fmt.Println(strings.Join(moves, " "))
// 		}
// 		if !stillMoving {
// 			finished = true
// 		}
// 	}
// }
package main

import (
	"fmt"
	"sort"
	"strings"
)

func ScheduleTrains(paths [][]string, trainCount int) {
	type Train struct {
		name     string
		path     []string
		delay    int
		position int
	}

	var trains []Train

	// Group paths by their length
	lengthGroups := make(map[int][][]string)
	var lengths []int

	for _, path := range paths {
		l := len(path)
		lengthGroups[l] = append(lengthGroups[l], path)
	}

	// Sort lengths ascending
	for l := range lengthGroups {
		lengths = append(lengths, l)
	}
	sort.Ints(lengths)

	// Choose the shortest group
	var usablePaths [][]string
	for _, l := range lengths {
		usablePaths = lengthGroups[l]
		if len(usablePaths) >= 1 {
			break
		}
	}

	// fallback: if somehow no usable paths (shouldn't happen), use all
	if len(usablePaths) == 0 {
		usablePaths = paths
	}

	// Assign trains using only usablePaths
	for i := 0; i < trainCount; i++ {
		path := usablePaths[i%len(usablePaths)]
		delay := i / len(usablePaths)
		trains = append(trains, Train{
			name:  fmt.Sprintf("T%d", i+1),
			path:  path[1:], // exclude start station
			delay: delay,
		})
	}

	// Debug train assignments
	fmt.Println("Train assignments:")
	for _, train := range trains {
		fmt.Printf("%s: %v (delay: %d)\n", train.name, train.path, train.delay)
	}
	fmt.Println("--------------------------------")

	// Simulation
	finished := false
	turn := 0

	for !finished {
		turn++
		moves := []string{}
		occupied := make(map[string]bool)
		stillMoving := false

		for i := range trains {
			train := &trains[i]
			if train.delay > 0 {
				train.delay--
				continue
			}
			if train.position >= len(train.path) {
				continue
			}
			next := train.path[train.position]
			if !occupied[next] {
				moves = append(moves, fmt.Sprintf("%s-%s", train.name, next))
				occupied[next] = true
				train.position++
				stillMoving = true
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
		if !stillMoving {
			finished = true
		}
	}
}
