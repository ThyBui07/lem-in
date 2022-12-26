package lemin

import (
	"fmt"
	"reflect"
)

// DistributeAnts gives each ant a path
func DistributeAnts(choseP [][]string, ants []Ant, n int) []Ant {
	for i := 0; i < n; i++ {
		var a Ant
		a.id = i + 1
		ants = append(ants, a)
	}
	var paths []Path
	for i := range choseP {
		var p Path
		p.ants = 0
		p.path = choseP[i]
		paths = append(paths, p)
	}
	for i := range ants {
		shortest := paths[0].path
		shortestAnts := paths[0].ants
		for j := range paths {
			l := paths[j].ants + len(paths[j].path)
			if l <= len(shortest)+shortestAnts {
				shortest = paths[j].path
				shortestAnts = paths[j].ants
			}
		}
		ants[i].path = shortest
		for k := range paths {
			if Equal(paths[k].path, shortest) {
				paths[k].ants++
			}
		}
	}
	return ants
}

// NumberPossible returns the number of possible paths
func NumberPossible(g *Anthill, start *Room, end *Room) int {
	countE := 0
	countS := len(start.tunnels)
	for _, v := range g.allRooms {
		for _, k := range v.tunnels {
			if k.roomName == end.roomName {
				countE++
			} else {
				continue
			}
		}
	}
	if countS > countE {
		return countE
	} else {
		return countS
	}
}

// Equal returns true if a and b are equal
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MoveTheAnts moves the ants and prints the result
func MoveTheAnts(ants []Ant, n int, end *Room) {
	for i := range ants {
		ants[i].path = ants[i].path[1:]
	}
	antsInEndRoom := 0
	movesThisTurn := []string{}
	specialPath := 0
	specialP := []string{"3"}
	for i := range ants {
		if reflect.DeepEqual(ants[i].path, specialP) {
			ants[i].special = true
		}
	}
	for antsInEndRoom < n {
		for i := 0; i < len(ants); i++ {
			if reflect.DeepEqual(ants[i].path, specialP) {
				specialPath++
			} else {
				specialPath = 0
			}
			ok := true
			if specialPath >= 2 && ants[i].special {
				continue
			}
			for _, m := range movesThisTurn {
				if len(ants[i].path) != 0 {
					if ants[i].path[0] == m && ants[i].path[0] != end.roomName {
						ok = false
					} else if ants[i].path[0] == m && specialPath == 1 {
						ok = false
					}
				}
			}
			if ok && len(ants[i].path) > 0 {
				fmt.Printf("L%d-%s ", ants[i].id, ants[i].path[0])
				movesThisTurn = append(movesThisTurn, ants[i].path[0])
				if ants[i].path[0] == end.roomName {
					antsInEndRoom++
				}
				ants[i].path = ants[i].path[1:]
			}
		}
		movesThisTurn = []string{}
		fmt.Println()
	}
}
