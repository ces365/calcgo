package main

import (
	"fmt"

	"gitlab.com/relnod/calcgo"
)

func main() {
	number := calcgo.Interpret("1 + 1")

	fmt.Println(number)

}