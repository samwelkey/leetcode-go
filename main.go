package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := strings.Split("hello world  ", " ")
	a := strings.Trim("hello world    ", " ")
	fmt.Println(a)
}
