package lemin

import "strings"

// formatElem removes comments and empty string before, between, and after elements and returns a the formatted slice of string
func formatElem(data []string) []string {
	var temp []string
	// store to temp only elements with no empty string & format element so that there are no spaces btw element
	for i := 0; i < len(data); i++ {
		if data[i] != "" {
			temp = append(temp, rmEmptySpaceBtwElem(data[i]))
		}
	}

	for i := 0; i < len(temp); i++ {
		if onlySpaces(temp[i]) {
			temp = append(temp[:i], temp[i+1:]...)
			i--
		}

		// rm #comments from slice
		if temp[i] == "##start" || temp[i] == "##end" {
			continue
		}
		if temp[i][0] == '#' {
			temp = append(data[:i], temp[i+1:]...)
			i--
		}
	}
	return temp
}

// onlySpaces checks that the string data contains only spaces and returns a bool
func onlySpaces(s string) bool {
	for _, c := range s {
		if c != 32 {
			return false
		}
	}
	return true
}

// rmEmptySpaceBtwElem removes empty spaces between a string and returns the formatted string
func rmEmptySpaceBtwElem(s string) string {
	var input []string
	var s2 string
	link := false
	if isLink(s) {
		link = true
	}
	inputData := strings.Split(s, " ")
	//delete all empty strings from slice
	for i := 0; i < len(inputData); i++ {
		if inputData[i] != "" {
			input = append(input, inputData[i])
		}
	}
	if link {
		s2 = strings.Join(input, "")
	} else {
		s2 = strings.Join(input, " ")
	}
	return strings.TrimSpace(s2)
}
