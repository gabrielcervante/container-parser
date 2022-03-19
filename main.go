package main

import "fmt"

func main() {

	fmt.Println(ContainerParser("{()}"))

}

func ContainerParser(container string) bool {

	containerSlice := []rune(container)

	var openedParentheses, closedParenthese, openedCurlybrace, closedCurlybrace, openedBracket, closedBracket int

	for i := range containerSlice {

		switch {
		case string(containerSlice[i]) == ")" && i == 0 || string(containerSlice[i]) == "(" && i == len(containerSlice)-1 || string(containerSlice[i]) == "]" && i == 0 || string(containerSlice[i]) == "[" && i == len(containerSlice)-1 || string(containerSlice[i]) == "}" && i == 0 || string(containerSlice[i]) == "{" && i == len(containerSlice)-1:
			return false
		case string(containerSlice[i]) == ")" && string(containerSlice[i-1]) == "[" || string(containerSlice[i]) == ")" && string(containerSlice[i-1]) == "{":
			return false
		case string(containerSlice[i]) == "]" && string(containerSlice[i-1]) == "(" || string(containerSlice[i]) == "]" && string(containerSlice[i-1]) == "{":
			return false
		case string(containerSlice[i]) == "}" && string(containerSlice[i-1]) == "[" || string(containerSlice[i]) == "}" && string(containerSlice[i-1]) == "(":
			return false
		case string(containerSlice[i]) == "(" && string(containerSlice[i+1]) == "]" || string(containerSlice[i]) == "(" && string(containerSlice[i+1]) == "}":
			return false
		case string(containerSlice[i]) == "[" && string(containerSlice[i+1]) == ")" || string(containerSlice[i]) == "[" && string(containerSlice[i+1]) == "}":
			return false
		case string(containerSlice[i]) == "{" && string(containerSlice[i+1]) == "]" || string(containerSlice[i]) == "{" && string(containerSlice[i+1]) == ")":
			return false
		case string(containerSlice[i]) == "(":
			openedParentheses++
		case string(containerSlice[i]) == ")":
			closedParenthese++
		case string(containerSlice[i]) == "{":
			openedCurlybrace++
		case string(containerSlice[i]) == "}":
			closedCurlybrace++
		case string(containerSlice[i]) == "[":
			openedBracket++
		case string(containerSlice[i]) == "]":
			closedBracket++
		}

	}

	if (openedParentheses+closedParenthese)%2 == 1 || (openedCurlybrace+closedCurlybrace)%2 == 1 || (openedBracket+closedBracket)%2 == 1 {
		return false
	}

	return true

}
