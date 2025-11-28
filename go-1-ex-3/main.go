package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln("the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln("the dice was rolled at", when)

	eyesFile, err := os.Create("eyes.txt")

	logFile, err := os.Create("dice.log")
	// TODO: how to write the output into eyes.txt and dice.log?
	// go run go-1-ex-3/main.go TODO
}
