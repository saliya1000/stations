package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Fprintln(os.Stderr, "Error: Too few arguments.")
		fmt.Println("Usage: go run . network.map start end numTrains")
		return
	} else if len(os.Args) > 5 {
		fmt.Fprintln(os.Stderr, "Error: Too many arguments.")
		fmt.Println("Usage: go run . network.map start end numTrains")
		return
	}

	filename := os.Args[1]
	start := os.Args[2]
	end := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("Invalid number of trains:", err)
		return
	}

	network, err := ParseNetworkFile(filename)
	if err != nil {
		fmt.Println("Failed to parse network:", err)
		return
	}

	result, err := ScheduleTrains(*network, start, end, numTrains)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Train movements:")
	for i, line := range result {
		fmt.Printf("Turn %d. %s\n", i+1, line)
	}
}
