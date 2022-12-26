package lemin

import (
	"fmt"
	"os"
	"strings"
)

type Anthill struct {
	numOfAnts int
	startRoom *Room
	endRoom   *Room
	allRooms  []*Room
	allLinks  []string
}

type Room struct {
	roomName string
	tunnels  []*Room
	visited  bool
}

type Ant struct {
	id      int
	path    []string
	special bool
}

type Path struct {
	path []string
	ants int
}

// AddRoom is a method of Anthill that adds room to AllRooms
func (g *Anthill) AddRoom(s string) {
	if contains(g.allRooms, s) {
		err := fmt.Errorf("Room %v not added because it is an existing RoomName", s)
		fmt.Println(err.Error())
	} else {
		g.allRooms = append(g.allRooms, &Room{roomName: s})
	}
}

// contains checks whether there is an existing equal room before adding and returns a bool
func contains(s []*Room, k string) bool {
	for _, v := range s {
		if k == v.roomName {
			return true
		}
	}
	return false
}

// GetRoom is a method of Anthill that returns the room with the given name or nil if not found
func (g *Anthill) GetRoom(k string) *Room {
	for i, v := range g.allRooms {
		if v.roomName == k {
			return g.allRooms[i]
		}
	}
	return nil
}

// CreateTunnels is a method of Anthill that creates links between the tunnels
func (g *Anthill) CreateTunnels() {
	for _, v := range g.endRoom.tunnels {
		g.AddTunnel(v.roomName, g.endRoom.roomName)
	}
	for i := 0; i < len(g.allLinks); i++ {
		temp := strings.Split(g.allLinks[i], "-")
		if temp[0] == temp[1] {
			fmt.Println("ERROR: invalid data format, room cannot link to iself")
			os.Exit(0)
		}
		g.AddTunnel(temp[0], temp[1])
	}
}

// AddTunnel is a method of Anthill that adds tunnels from and to each room
func (g *Anthill) AddTunnel(from, to string) {
	//get Room
	fromRoom := g.GetRoom(from)
	toRoom := g.GetRoom(to)

	//check error
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("invalid edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.tunnels, to) {
		err := fmt.Errorf("existing edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add Tunnel
		fromRoom.tunnels = append(fromRoom.tunnels, toRoom)
		toRoom.tunnels = append(toRoom.tunnels, fromRoom)
	}
}

// CheckPathToEndFromStartRoom is a method of Anthill that checks if there is a path from start to end room
func (g *Anthill) CheckPathToEndFromStartRoom() {
	countStart := 0
	countEnd := 0
	var linksStart []string
	var linksEnd []string
	var colonIndex int
	for _, c := range g.allLinks {
		colonIndex = strings.Index(c, "-")
		linksStart = append(linksStart, c[:colonIndex])
		linksEnd = append(linksEnd, c[colonIndex+1:])
	}
	for i := 0; i < len(linksStart); i++ {
		if linksStart[i] == g.startRoom.roomName {
			countStart++
		}
	}
	for i := 0; i < len(linksEnd); i++ {
		if linksEnd[i] == g.endRoom.roomName {
			countEnd++
		}
	}
	if countStart <= 0 {
		fmt.Println("ERROR: invalid data format, no path from start")
		os.Exit(0)
	}
	if countEnd <= 0 {
		fmt.Println("ERROR: invalid data format, no path to end")
		os.Exit(0)
	}
}
