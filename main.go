package main

import (
	"fmt"
	"github.com/marinco/gh-actions-go/cmd"
)

func main() {
	name := "World"
	fmt.Println(cmd.Hello(name))
}
