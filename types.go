package main

type Network struct {
	Stations map[string][]string
}

func (n *Network) GetStationByName(name string) *string {
	if _, exists := n.Stations[name]; exists {
		return &name
	}
	return nil
}

func (n *Network) GetConnectedStations(name string) []string {
	return n.Stations[name]
}

type Train struct {
	ID        string
	Path      []string
	EntryTurn int
}

type PathAssignment struct {
	PathIndex int
	EntryTurn int
}
