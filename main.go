package main

import (
	
	"fmt"
	"os"
	"strconv"
	

)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error: incorrect number of arguments")
		os.Exit(1)
	}

	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	trainCount, err := strconv.Atoi(os.Args[4])
	if err != nil || trainCount <= 0 {
		fmt.Fprintln(os.Stderr, "Error: number of trains is not a valid positive integer")
		os.Exit(1)
	}

	stations, connections, err := ParseMap(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	

	graph := BuildGraph(stations, connections)
	//fmt.Println(graph)
	if _, ok := stations[startStation]; !ok {
		fmt.Fprintln(os.Stderr, "Error: start station does not exist")
		os.Exit(1)
	}
	if _, ok := stations[endStation]; !ok {
		fmt.Fprintln(os.Stderr, "Error: end station does not exist")
		os.Exit(1)
	}
	if startStation == endStation {
		fmt.Fprintln(os.Stderr, "Error: start and end station are the same")
		os.Exit(1)
	}

	
	paths := FindAllPaths(graph, startStation, endStation, 20) 

	fmt.Println("Number of shortest paths:", len(paths))
	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no path between the start and end stations")
		os.Exit(1)
	}

	ScheduleTrains(paths, trainCount)
	
}

