 # ✨Train Scheduler Test Cases✨

This project simulates train scheduling on a given network.
It takes a network map file, a start node, an end node, and the number of trains as input.
The program then calculates and displays a turn-by-turn schedule of train movements to efficiently move all trains from the start node to the end node.
This Folder used predefined test cases to indentify program usability and felxibility.

## Test Usage

To run the train scheduler test case, use the following command:

```bash
go run . test/network.map start_node end_node num_trains
```

For example:

```bash
go run . test/one.map two six 6
```

## File Descriptions

- `duplicate_stations.go`: This file used to indentify duplicate Stations. used albinoni station name twice
- `not_exist.map`: This file used to indetify not exist station find. part station is missing.
- `same_coordinates.go`: This file used same coordinates for different stations.
- `station_invalid.go`:This file used to indentify station name invalid. avoid characters symbols


 ## 📦Dependencies

- Go 1.23.0 (or compatible)

✨ This project is designed to be simple, efficient, and easy to use. Contributions and feedback are always welcome! 🚀