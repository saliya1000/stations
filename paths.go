package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxStations = 10000 // Set the limit for stations

func ParseNetworkFile(filename string) (*Network, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	network := &Network{Stations: make(map[string][]string)}
	scanner := bufio.NewScanner(file)
	isConnection := false
	stationsCount := 0                         // Keep track of the number of stations
	usedCoordinates := make(map[string]string) // key: "x,y", value: station name
	connectionSet := make(map[string]bool)     // key: "A-B" or "B-A" (sorted)
	foundStationSection := false
	foundConnectionSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "stations:" {
			isConnection = false
			foundStationSection = true
			continue
		}
		if line == "connections:" {
			isConnection = true
			foundConnectionSection = true
			continue
		}

		if !isConnection {
			parts := strings.Split(line, ",")
			if len(parts) >= 3 {
				name := strings.TrimSpace(parts[0])
				xStr := strings.TrimSpace(parts[1])
				yStr := strings.TrimSpace(parts[2])

				x, err1 := strconv.Atoi(xStr)
				y, err2 := strconv.Atoi(yStr)

				if err1 != nil || err2 != nil || x < 0 || y < 0 {
					return nil, fmt.Errorf("invalid station coordinates: %s", line)

				}
				// check if the coordinates are already used
				coordKey := fmt.Sprintf("%d,%d", x, y)
				if existingStation, exists := usedCoordinates[coordKey]; exists {
					return nil, fmt.Errorf("duplicate coordinates (%d,%d) found for stations '%s' and '%s'", x, y, existingStation, name)
				}
				usedCoordinates[coordKey] = name

				// Check if we've reached the station limit
				if stationsCount >= MaxStations {
					return nil, fmt.Errorf("station limit exceeded, cannot add more than %d stations", MaxStations)
				}

				network.Stations[name] = []string{}
				stationsCount++ // Increment the station count
			}
		} else {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				a := strings.TrimSpace(parts[0])
				b := strings.TrimSpace(parts[1])

				// Validate that both stations exist
				if _, ok := network.Stations[a]; !ok {
					return nil, fmt.Errorf("connection error: station '%s' does not exist", a)
				}
				if _, ok := network.Stations[b]; !ok {
					return nil, fmt.Errorf("connection error: station '%s' does not exist", b)
				}

				keyA := a
				keyB := b
				if keyA > keyB {
					keyA, keyB = keyB, keyA
				}
				connectionKey := keyA + "-" + keyB

				// Check for duplicate connection
				if connectionSet[connectionKey] {
					return nil, fmt.Errorf("duplicate connection found between '%s' and '%s'", a, b)
				}
				connectionSet[connectionKey] = true

				// Add the connection in both directions
				network.Stations[a] = append(network.Stations[a], b)
				network.Stations[b] = append(network.Stations[b], a)
			}
		}
	}
	if !foundStationSection {
		//fmt.Printf("Error: Missing stations section in the file")
		return nil, fmt.Errorf("Error: missing stations section in the file")
	}
	if !foundConnectionSection {
		//fmt.Printf("Error: Missing Connection section in the file")
		return nil, fmt.Errorf("Error: missing Connections section in the file")
	}

	return network, scanner.Err()
}

func FindPaths(network *Network, start, end string) ([][]string, error) {
	if err := ValidateStationExists(network, start, true); err != nil {
		return nil, err
	}
	if err := ValidateStationExists(network, end, false); err != nil {
		return nil, err
	}

	var paths [][]string
	visited := make(map[string]bool)
	var currentPath []string

	var dfs func(station string)
	dfs = func(station string) {
		visited[station] = true
		currentPath = append(currentPath, station)

		if station == end {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			for _, next := range network.GetConnectedStations(station) {
				if !visited[next] {
					dfs(next)
				}
			}
		}

		visited[station] = false
		currentPath = currentPath[:len(currentPath)-1]
	}

	dfs(start)
	fmt.Println("All paths found:")
	for i, path := range paths {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
	fmt.Println("-------------------------------")

	return paths, nil
}
