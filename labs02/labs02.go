package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello,world!"
	s2 := "llo"
	fmt.Println(strings.Contains(s1, s2))
}
