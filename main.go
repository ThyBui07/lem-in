package main

import (
	"fmt"
	allFunc "lemin/functions"
)

func main() {
	// check validity of filename
	args := allFunc.ValidArgument()
	if args == nil {
		return
	}
	// create Anthill and rooms
	myAnthill, err := allFunc.CreateAntHill("./audit/" + args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	myAnthill.CheckPathToEndFromStartRoom() // check if there is path from start and path to end
	myAnthill.CreateTunnels()               // create tunnels
	myAnthill.PrintResult(args[0])          // print result
}
