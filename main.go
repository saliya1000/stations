package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Fprintln(os.Stderr, "error: too few arguments")
		fmt.Println("Usage: go run . network.map start end numTrains")
		return
	} else if len(os.Args) > 5 {
		fmt.Fprintln(os.Stderr, "error: too many arguments.")
		fmt.Println("Usage: go run . network.map start end numTrains")
		return
	}

	filename := os.Args[1]
	start := os.Args[2]
	end := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Fprintln(os.Stderr,"error: Invalid number of trains:", err)
		return
	}

	network, err := ParseNetworkFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr,"error: failed to parse network:", err)
		return
	}

	if numTrains <= 0 {
		fmt.Fprintln(os.Stderr, "Error: Number of trains must be a positive integer.")
		return
	}

	

	if start == end {
		fmt.Fprintln(os.Stderr, "Error: Start and End stations are Same.")
		return
	}

	result, err := ScheduleTrains(*network, start, end, numTrains)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Train movements:")
	for i, line := range result {
		fmt.Printf("Turn %d. %s\n", i+1, line)
	}
}
