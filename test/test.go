package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`([a-zA-Z]+)=([^=]*\w)(?:\s|$)`)

	var input = "filter  lo=jakarta, helo type=fullstack"

	var g = regex.FindAllString(input, -1)
	var strs = input[:8-1]

	fmt.Println(strs)
	fmt.Println(g[0])

}
