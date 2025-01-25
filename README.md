# Lem-in: A Digital Ant Farm

## Overview

This project is a digital version of an ant farm. The goal is to create a program, `lem-in`, that simulates ants finding the quickest way across a colony of rooms and tunnels. The program reads a file describing the ants and the colony and outputs the simulation of the ants' movements.

## How It Works

1. **Ant Farm Setup:**
   - The colony is composed of rooms and tunnels.
   - Ants start in the `##start` room and must find their way to the `##end` room.
   - The program determines the quickest way to move all ants from `##start` to `##end` using the minimum number of moves.

2. **Input File:**
   - The input file specifies the number of ants, room details, and tunnels.
   - Example file:
     ```
     4
     ##start
     0 0 3
     2 2 5
     3 4 0
     ##end
     1 8 3
     0-2
     2-3
     3-1
     ```

3. **Simulation Output:**
   - The output shows the content of the input file and the step-by-step movement of ants:
     ```
     $ go run . example00.txt
     4
     ##start
     0 0 3
     2 2 5
     3 4 0
     ##end
     1 8 3
     0-2
     2-3
     3-1

     L1-2
     L1-3 L2-2
     L1-1 L2-3 L3-2
     L2-1 L3-3 L4-2
     L3-1 L4-3
     L4-1
     $
     ```

## Features

- **Shortest Path Calculation:** 
  - Finds the optimal way to move all ants in the fewest moves.
  - Considers complex colonies with multiple paths and weights.

- **Input Validation:**
  - Detects and returns specific error messages for invalid input, such as:
    - No `##start` or `##end` room.
    - Duplicate rooms or invalid room coordinates.
    - Links to unknown rooms.
    - Invalid number of ants.
  - Example error messages:
    - `ERROR: invalid data format, no start room found`
    - `ERROR: invalid data format, invalid number of Ants`

- **Robust Simulation:**
  - Handles edge cases like:
    - Colonies with no path from `##start` to `##end`.
    - Rooms linked to themselves, causing infinite loops.

## Usage
   ```bash
   go run . <input_file>
   ```
   Example : ```go run . example00.txt```

## Key considerations

- The shortest path is not always the simplest.
- Some colonies might not have a path from ##start to ##end.
- The program will exit gracefully in cases of invalid input.

## Conclusion 

Lem-in is a robust program for simulating the most efficient movements of ants across a digital colony. It demonstrates advanced pathfinding and validation techniques while handling complex scenarios and edge cases. 
