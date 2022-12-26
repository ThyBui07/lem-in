package lemin

import (
	"fmt"
	"os"
)

// PrintResult reads the ant.txt and print the data, find all Paths, find best paths, distribute the ants, and move the ants
func (g *Anthill) PrintResult(filename string) {
	fileBytes, err := os.ReadFile("./audit/" + filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileBytes))
	fmt.Println()
	allPaths := FindAllPaths(g.startRoom, g.endRoom)
	var ants []Ant
	FindBestPaths(allPaths)
	maxNumberOfPaths := NumberPossible(g, g.startRoom, g.endRoom)
	if len(combinePaths) <= maxNumberOfPaths {
		for _, v := range combinePaths {
			if len(v) == maxNumberOfPaths {
				ants = DistributeAnts(v, ants, g.numOfAnts)
				MoveTheAnts(ants, g.numOfAnts, g.endRoom)
			}
		}
	} else {
		CalculateSteps(combinePaths)
		ants = DistributeAnts(Chosen, ants, g.numOfAnts)
		MoveTheAnts(ants, g.numOfAnts, g.endRoom)
	}
}
