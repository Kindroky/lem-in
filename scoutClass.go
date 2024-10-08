package main

import "fmt"

type Scout struct {
	CurrentRoom   *Piece
	MemorizedPath []*Piece
}

// Explore method for Scouts
func (s *Scout) Explore(allPaths *[][]*Piece) {
	// Mark the current room as visited
	s.CurrentRoom.Visited = true

	s.MemorizedPath = append(s.MemorizedPath, s.CurrentRoom)

	// Check if the current room is the end room
	if s.CurrentRoom.End {
		// Save the current path
		*allPaths = append(*allPaths, append([]*Piece{}, s.MemorizedPath...))
	}

	// Explore adjacent rooms
	if !s.CurrentRoom.End {
		for _, nextRoom := range s.CurrentRoom.GetSalles() {
			// Only explore if the next room has not been visited
			if !nextRoom.Visited {
				scout := &Scout{
					CurrentRoom:   nextRoom,
					MemorizedPath: s.MemorizedPath,
				}
				scout.Explore(allPaths) // Recur to explore the next room
			}
		}
	}
	s.CurrentRoom.Visited = false
}

// Function to start the exploration with multiple scouts
func StartExploration(startRoom *Piece) [][]*Piece {
	allPaths := [][]*Piece{}
	scout := &Scout{CurrentRoom: startRoom}
	scout.Explore(&allPaths)
	return allPaths
}

func SortPath(allPaths [][][]*Piece) [][][]*Piece {
	for range allPaths {
		for j := range allPaths {
			for i := len(allPaths) - 1; i > j; i-- {
				if len(allPaths[j]) < len(allPaths[i]) {
					allPaths[j], allPaths[i] = allPaths[i], allPaths[j]
				}
			}
		}
	}
	return allPaths
}

func IndiePaths(allPaths [][]*Piece) [][][]*Piece {
	independentGroups := [][][]*Piece{}

	for _, path := range allPaths {
		nouvGroup := [][]*Piece{path}
		for _, indePath := range allPaths {
			valid := true
			for _, groupPath := range nouvGroup {
				if !IsInde(indePath, groupPath) {
					valid = false
				}
			}
			if valid {
				nouvGroup = append(nouvGroup, indePath)
			}
		}
		independentGroups = append(independentGroups, nouvGroup)
	}
	return independentGroups
}

func PrintGroups(groups [][][]*Piece) {
	for _, g := range groups {
		for _, l := range g {
			for _, s := range l {
				fmt.Print(s.Nom)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func IsInde(path, path2 []*Piece) bool {
	for _, s := range path {
		if !s.Start && !s.End {
			for _, c := range path2 {
				if s.Nom == c.Nom {
					return false
				}
			}
		}
	}
	return true
}

func FindGroupsWithMostPaths(groups [][][]*Piece) [][][]*Piece {
	var mostIndiePaths [][][]*Piece //initialize a slice of groups with most indie paths
	maxPaths := 0

	for _, group := range groups {
		if len(group) > maxPaths {
			maxPaths = len(group)
			mostIndiePaths = [][][]*Piece{group}
		} else if len(group) == maxPaths {
			mostIndiePaths = append(mostIndiePaths, group)
		}
	}
	return mostIndiePaths
}

func FindShortestPath(groups [][][]*Piece) [][]*Piece {
	if len(groups) == 1 { // If there is only 1 group, return it
		return groups[0]
	}

	minRooms := -1
	bestGroupIndex := -1

	for i := 0; i < len(groups); i++ {
		nbOfRooms := 0
		for j := 0; j < len(groups[i]); j++ {
			nbOfRooms += len(groups[i][j]) // Total number of rooms in each group
		}

		// Update the best group if this group has fewer rooms
		if minRooms == -1 || nbOfRooms < minRooms {
			minRooms = nbOfRooms
			bestGroupIndex = i
		} else if nbOfRooms == minRooms {
			// Find the shortest path length in the current best group
			bestPathLength := -1
			for _, path := range groups[bestGroupIndex] {
				if bestPathLength == -1 || len(path) < bestPathLength {
					bestPathLength = len(path)
				}
			}

			// Check the current group for a shorter path
			for _, path := range groups[i] {
				if len(path) < bestPathLength {
					bestGroupIndex = i // Update the current group to the bestgroup if it contains a shorter path
				}
			}
		}
	}

	return groups[bestGroupIndex] // Return the best group
}

/*
independentGroups := [][][]*Piece{} // Liste de groupes de chemins indépendants

	for _, path := range allPaths {
		wasAddedToGroup := false
		newGroups := [][][]*Piece{} // Garde une trace des nouveaux groupes qui incluront le chemin

		// Parcourir tous les groupes existants
		for _, group := range independentGroups {
			conflict := false

			// Vérification des conflits de salles intermédiaires avec les chemins du groupe
			for _, otherPath := range group {
				for i := 1; i < len(path)-1; i++ { // Ignorer la première et la dernière pièce
					for j := 1; j < len(otherPath)-1; j++ {
						if path[i].Nom == otherPath[j].Nom {
							conflict = true
							break
						}
					}
					if conflict {
						break
					}
				}
				if conflict {
					break
				}
			}

			// Si aucun conflit, ajouter le chemin à ce groupe
			if !conflict {
				group = append(group, path)
				newGroups = append(newGroups, group)
				wasAddedToGroup = true
			}
		}

		// Si le chemin peut être ajouté à plusieurs groupes, les nouvelles versions de ces groupes sont ajoutées
		if wasAddedToGroup {
			independentGroups = append(independentGroups, newGroups...)
		}

		// Si le chemin ne peut pas être ajouté à un groupe existant, on crée un nouveau groupe
		if !wasAddedToGroup {
			independentGroups = append(independentGroups, [][]*Piece{path})
		}
	} */
