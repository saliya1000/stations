package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Station struct {
	Name string
	X    int
	Y    int
}

func ParseMap(filePath string) (map[string]Station, [][2]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file")
	}
	defer file.Close()

	stations := make(map[string]Station)
	coords := make(map[string]bool)
	connections := [][2]string{}
	
	scanner := bufio.NewScanner(file)
	mode := ""
	nameRe := regexp.MustCompile(`^[a-z0-9_]+$`)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "#")[0]
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "stations:" {
			mode = "stations"
			continue
		} else if line == "connections:" {
			mode = "connections"
			continue
		}

		switch mode {
		case "stations":
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, nil, fmt.Errorf("invalid station entry")
			}
			name := strings.TrimSpace(parts[0])
			x, err1 := strconv.Atoi(strings.TrimSpace(parts[1]))
			y, err2 := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err1 != nil || err2 != nil || x < 0 || y < 0 {
				return nil, nil, fmt.Errorf("invalid coordinates")
			}
			if !nameRe.MatchString(name) {
				return nil, nil, fmt.Errorf("invalid station name")
			}
			coordKey := fmt.Sprintf("%d,%d", x, y)
			if coords[coordKey] {
				return nil, nil, fmt.Errorf("duplicate coordinates")
			}
			if _, exists := stations[name]; exists {
				return nil, nil, fmt.Errorf("duplicate station name")
			}
			coords[coordKey] = true
			stations[name] = Station{name, x, y}

		case "connections":
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("invalid connection entry")
			}
			a := strings.TrimSpace(parts[0])
			b := strings.TrimSpace(parts[1])
			if a == b {
				return nil, nil, fmt.Errorf("connection loops to itself")
			}
			if _, ok := stations[a]; !ok {
				return nil, nil, fmt.Errorf("station %s not found", a)
			}
			if _, ok := stations[b]; !ok {
				return nil, nil, fmt.Errorf("station %s not found", b)
			}
			for _, c := range connections {
				if (c[0] == a && c[1] == b) || (c[0] == b && c[1] == a) {
					return nil, nil, fmt.Errorf("duplicate connection")
				}
			}
			connections = append(connections, [2]string{a, b})
		default:
			continue
		}
	}

	return stations, connections, nil
}