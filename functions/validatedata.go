package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isValidDataFormat checks that input file is valid
func isValidDataFormat(sl []string) bool {
	if !validEndStart(sl) {
		return false
	}
	if !containsLinks(sl) {
		return false
	}
	if !isComment(sl) {
		return false
	}
	return true
}

// validEndStart checks validity of ##start, ##end, and their respective rooms and returns a bool
func validEndStart(sl []string) bool {
	end := false
	start := false
	for i := 0; i < len(sl); i++ {
		if sl[i] == "##start" {
			start = true
			if !isRoom(sl[i+1]) {
				fmt.Println("ERROR: invalid data format, no room found after ##start")
				return false
			}
		} else if sl[i] == "##end" {
			end = true
			if !isRoom(sl[i+1]) {
				fmt.Println("ERROR: invalid data format, no room found after ##end")
				return false
			}
		}
		// check that ##start & ##end are not at EOF
		if sl[len(sl)-1] == "##start" || sl[len(sl)-1] == "##end" {
			fmt.Println("ERROR: invalid data format, ##start and ##end cannot be at EOF")
			return false
		}

	}
	if !end && !start {
		fmt.Println("ERROR: invalid data format, no ##start and ##end found")
		return false
	} else if !start {
		fmt.Println("ERROR: invalid data format, no ##start found")
		return false
	} else if !end {
		fmt.Println("ERROR: invalid data format, no ##end found")
		return false
	}
	return true
}

// containsLinks bool checks that the slice of string contains links and returns a bool
func containsLinks(sl []string) bool {
	containsLinks := false
	for i := 0; i < len(sl); i++ {
		if isLink(sl[i]) {
			containsLinks = true
			continue
		}
	}
	if !containsLinks {
		fmt.Println("ERROR: invalid data format, no links found")
		return false
	}
	return true
}

// isLink checks that the slice is a link and returns a bool
func isLink(s string) bool {
	s = strings.TrimSpace(s)                // remove trailing spaces
	s = strings.Join(strings.Fields(s), "") // remove spaces in between
	return len(strings.Split(s, "-")) == 2
}

// isComment checks that the slice is a comment and returns a bool
func isComment(sl []string) bool {
	for i := 0; i < len(sl); i++ {
		if string(sl[i][0]) == "#" {
			return true
		}
	}
	return false
}

// isNum checks that the string data is a number and returns a bool
func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// isRoom checks that the string data is a room and returns a bool
func isRoom(s string) bool {
	if s[0] == 'L' {
		fmt.Println("ERROR: invalid data format, invalid room name")
		os.Exit(0)
	}
	input := strings.Fields(s)
	if len(input) != 3 {
		return false
	}
	if !isNum(input[1]) || !isNum(input[2]) {
		fmt.Println("ERROR: invalid data format, invalid room coordinates")
		os.Exit(0)
	}
	return true
}

// ValidArgument checks validity of args and returns the args
func ValidArgument() []string {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: no file provided")
		return nil
	}

	if args[0] != "badexample00.txt" && args[0] != "badexample01.txt" && args[0] != "example00.txt" && args[0] != "example01.txt" && args[0] != "example02.txt" && args[0] != "example03.txt" && args[0] != "example04.txt" && args[0] != "example05.txt" && args[0] != "example06.txt" && args[0] != "example07.txt" {
		fmt.Println("ERROR: invalid file name")
		return nil
	}
	return args
}
