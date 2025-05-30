 # ✨Train Scheduler ✨

This project simulates train scheduling on a given network.
It takes a network map file, a start node, an end node, and the number of trains as input.
The program then calculates and displays a turn-by-turn schedule of train movements to efficiently move all trains from the start node to the end node.

## Usage

To run the train scheduler, use the following command:

```bash
go run . network.map start_node end_node num_trains
```

For example:

```bash
go run . network.map A D 3
```

## File Descriptions

- `main.go`: The main entry point for the application. Parses command-line arguments and orchestrates the simulation.
- `network.map`: An example file representing the train network. This file defines the nodes and connections in the network. (Note: The format of this file would need to be defined for users to create their own).
- `best_path.go`: Likely contains logic for finding the best path between two nodes in the network.
- `paths.go`: May contain data structures or functions related to handling multiple paths or sequences of paths.
- `simulate.go`: Contains the core logic for simulating the train movements and generating the schedule.
- `types.go`: Defines custom data types and structures used throughout the project (e.g., for representing the network, trains, paths).
- `errors.go`: Defines custom error types for the application.
- `go.mod`: The Go module file, specifying the module name and dependencies.

 ## 📦Dependencies

- Go 1.23.0 (or compatible)

✨ This project is designed to be simple, efficient, and easy to use. Contributions and feedback are always welcome! 🚀