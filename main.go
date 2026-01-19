package main

import "fmt"

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	newstr := str[len(str)-len(ending):]
	if newstr == ending {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Print(solution("", " "))
}
