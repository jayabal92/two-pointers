package main

import "fmt"

func isSunSequence(s string, t string) bool {
	i := 0
	j := 0

	for i < len(s) && j < len(t) {

		if s[i] == t[j] {
			i++
		}
		j++

	}

	return i == len(s)
}

func main() {
	t := "abcde"
	subSeq := "aced"
	fmt.Printf("%#v", isSunSequence(subSeq, t))
}
