package lemin

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errNumOfAnts = errors.New("ERROR: invalid data format, invalid number of ants")

// CreateAntHill creates the anthill and rooms
func CreateAntHill(filename string) (*Anthill, error) {
	data := getInputData(filename)
	data = formatElem(data)
	var numAnts int
	var allRooms []*Room
	myAnthill := Anthill{
		numOfAnts: numAnts,
		allRooms:  allRooms,
	}

	if isValidDataFormat(data) {
		var temp []string
		for i, line := range data {
			if i == 0 {
				// check num of ants
				numOfAnts, err := strconv.Atoi(data[0])
				if err != nil || numOfAnts < 1 {
					return nil, errNumOfAnts
				}
				myAnthill.numOfAnts = numOfAnts
			} else if isRoom(line) {
				temp = strings.Split(line, " ")
				myAnthill.AddRoom(temp[0])
				if data[i-1] == "##start" {
					myAnthill.startRoom = myAnthill.allRooms[len(myAnthill.allRooms)-1]
				}
				if data[i-1] == "##end" {
					myAnthill.endRoom = myAnthill.allRooms[len(myAnthill.allRooms)-1]
				}
			} else if isLink(line) {
				myAnthill.allLinks = append(myAnthill.allLinks, line)
			} else if line[0] != '#' {
				fmt.Println("ERROR: invalid data input")
				os.Exit(0)
			}
		}
	} else {
		os.Exit(0)
	}
	return &myAnthill, nil
}

// getInputData gets the input data from the file
func getInputData(filename string) []string {
	src, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(src) == 0 {
		fmt.Println("ERROR: file is empty")
		os.Exit(0)
	}
	data := strings.Split(strings.ReplaceAll(string(src), "\r\n", "\n"), "\n")
	return data
}
