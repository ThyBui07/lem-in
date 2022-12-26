package lemin

import (
	"sort"
)

var currentPath []string
var allPaths [][]string
var combinePaths [][][]string

// FindAllPaths takes in a start room and end room as parameters and returns a [][]string of all paths.
// It starts at the start room and appends the next path recursively till the end room
// based on whether the room has been visited, which is a bool, and that start room is not end room.
func FindAllPaths(start *Room, end *Room) [][]string {
	if start.visited {
		return allPaths
	}
	start.visited = true
	currentPath = append(currentPath, start.roomName)
	if start.roomName == end.roomName {
		temp := make([]string, len(currentPath))
		copy(temp, currentPath)
		allPaths = append(allPaths, temp)
		start.visited = false
		//backtracking- remove last element
		currentPath = currentPath[:len(currentPath)-1]
		return allPaths
	} else {
		for _, v := range start.tunnels {
			FindAllPaths(v, end)
		}
	}
	//backtracking- remove last element
	currentPath = currentPath[:len(currentPath)-1]
	start.visited = false
	return allPaths
}

// FindBestPaths appends paths that do not have intersections and stores to BestPaths
func FindBestPaths(allPaths [][]string) {
	var combinePath [][]string
	for i := 0; i < len(allPaths); i++ {
		combinePath = append(combinePath, allPaths[i])
		for j := i + 1; j < len(allPaths); j++ {
			var check bool
			for _, v := range combinePath {
				check = SlicesEqual(v[1:len(v)-1], allPaths[j][1:len(allPaths[j])-1])
				if check {
					break
				}
			}
			if !check {
				combinePath = append(combinePath, allPaths[j])
			} else {
				continue
			}
		}
		combinePaths = append(combinePaths, combinePath)
		combinePath = nil
	}
}

// SlicesEqual checks if two slices have an element with the same value with the same index
func SlicesEqual(a, b []string) bool {
	var equal bool
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				equal = true
				break
			}
		}
	}
	return equal
}

var flow float64
var m = PathCombos{}
var Chosen [][]string

type PathCombos struct {
	steps float64
	combo [][]string
}

// CalculateSteps calculates the number of steps for each path combination
func CalculateSteps(best [][][]string) {
	myPathCombos := []PathCombos{}
	for _, combo := range best {
		m.combo = combo
		numberofPaths := float64(len(combo))
		numberOfRooms := float64(0)
		for _, v := range combo {
			numberOfRooms += float64(len(v[1 : len(v)-1]))
		}
		flow = (numberOfRooms / numberofPaths) / numberofPaths
		m.steps = flow
		myPathCombos = append(myPathCombos, m)
	}
	keys := make([]float64, 0, len(myPathCombos))
	for _, v := range myPathCombos {
		keys = append(keys, v.steps)
	}
	sort.Float64s(keys)
	for _, v := range myPathCombos {
		if v.steps == keys[0] {
			Chosen = v.combo
			break
		}
	}
}
