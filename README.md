# Train Scheduler Project Overview

## 🎯 Project Purpose
The Train Scheduler is a Go application that simulates efficient train scheduling on a railway network. Given a network topology, starting station, destination station, and number of trains, it calculates the optimal turn-by-turn movement schedule to move all trains from start to destination in minimum time.

## 🏗️ Architecture Overview

### Core Components
1. **Network Parser** (`paths.go`) - Reads and validates network topology
2. **Path Finder** (`paths.go`) - Discovers all possible routes
3. **Path Optimizer** (`best_path.go`) - Selects optimal disjoint paths
4. **Scheduler** (`simulate.go`) - Creates train assignments and movements
5. **Main Controller** (`main.go`) - Orchestrates the entire process

## 📊 Data Flow

```
Network File → Parse Network → Find All Paths → Filter Disjoint Paths → 
Calculate Optimal Turns → Assign Trains → Simulate Movements → Output Schedule
```

## 🔧 Key Functions Breakdown

### 1. Network Parsing (`ParseNetworkFile`)
**Purpose**: Reads network topology from file and validates structure

**Process**:
- Parses station definitions with coordinates
- Validates station names (alphanumeric + underscore only)
- Checks for duplicate stations and coordinates
- Parses bidirectional connections
- Enforces maximum station limit (10,000)

**Validation Rules**:
- Station names must be unique and follow naming convention
- Coordinates must be unique positive integers
- Connections must reference existing stations
- No duplicate connections allowed

### 2. Path Discovery (`FindPaths`)
**Purpose**: Uses Depth-First Search to find all possible paths between start and end stations

**Algorithm**:
- DFS with backtracking to explore all routes
- Maintains visited set to prevent cycles
- Recursively explores all connected stations
- Returns all valid paths from start to destination

### 3. Path Optimization (`FilterDisjointPaths`)
**Purpose**: Selects the best combination of non-conflicting paths

**Strategy**:
- Identifies paths that don't share intermediate stations
- Uses brute-force combination testing
- Evaluates each combination using `CalculateMinTurns`
- Returns the combination that minimizes total turns

**Disjoint Logic**: Two paths are disjoint if they don't share any intermediate stations (start/end stations can be shared)

### 4. Turn Calculation (`CalculateMinTurns`)
**Purpose**: Determines minimum turns needed to move all trains

**Formula**:
- For each path of length L, trains can enter from turn 1 to turn (T-L)
- Total capacity = Σ(max(0, T-L)) for all paths
- Find minimum T where total capacity ≥ number of trains

### 5. Train Assignment (`CreateTrainAssignments`)
**Purpose**: Assigns specific trains to paths and entry turns

**Process**:
- Generates all possible (path, entry_turn) combinations
- Sorts by entry turn, then by path index
- Assigns first N combinations to N trains
- Ensures optimal utilization of available paths

### 6. Movement Simulation (`SimulateTrainMovements`)
**Purpose**: Generates turn-by-turn movement schedule

**Process**:
- For each turn, calculates active train positions
- Trains move one station per turn along their assigned path
- Outputs movements in format: "T1-StationB T2-StationC"
- Sorts movements alphabetically for consistent output

## 📋 Input/Output Format

### Input File Structure
```
stations:
waterloo,3,1
victoria,6,7
euston,11,23

connections:
waterloo-victoria
waterloo-euston
victoria-euston
```

### Command Line Usage
```bash
go run . network.map start_station end_station num_trains
```

### Command Line Usage for Test Files
```bash
go run . test/network.map start_station end_station num_trains
```

### Output Format
```
Turn 1. T1-victoria T2-euston
Turn 2. T1-euston T2-victoria
Turn 3. T1-st_pancras
```

## 🎛️ Algorithm Complexity

- **Path Finding**: O(V + E) * P where P is number of paths
- **Path Optimization**: O(2^P) where P is number of paths (exponential)
- **Turn Calculation**: O(T * P) where T is turns, P is paths
- **Simulation**: O(T * N) where T is turns, N is trains

## 🔍 Key Design Decisions

1. **Disjoint Path Strategy**: Prevents train collisions by ensuring no shared intermediate stations
2. **Greedy Assignment**: Assigns trains to earliest available slots for optimal throughput
3. **Exhaustive Path Search**: Finds all possible routes to ensure optimal solution
4. **Turn-based Simulation**: Simplifies scheduling by using discrete time steps

## 🚀 Optimization Opportunities

1. **Path Selection**: Current brute-force approach could use dynamic programming
2. **Memory Usage**: Large networks might benefit from streaming approaches
3. **Parallel Processing**: Path finding could be parallelized
4. **Heuristics**: Could add A* or other informed search algorithms

## 🎯 Use Cases

- Railway network optimization
- Resource scheduling simulation
- Graph traversal optimization
- Educational tool for algorithm demonstration

## 🔧 Error Handling

The system includes comprehensive validation:
- Invalid network files
- Non-existent stations
- Invalid train counts
- Unreachable destinations
- Duplicate network elements

This design ensures robust operation while providing clear error messages for debugging and user guidance.