package main

import (
	"fmt"
)

const (
	farmUrl = "localhost:1633"
	id      = "/:testminer/"
)

func main() {
	fmt.Println("Connecting to", farmUrl+id)

	return
}
