package main

import (
	"fmt"
	"os"

	common "github.com/halamix2/advent_of_code/cmd/01"
)

func main() {
	// rotationHints, err := common.LoadInput("cmd/01/in.txt")
	rotationHints, err := common.LoadInput("cmd/01/input.txt")
	if err != nil {
		fmt.Printf("Couldn't load hints: %v\n", err)
		os.Exit(1)
	}

	zeroCrossings := 0
	sm := common.SafeMachine{Position: 50}
	for _, rh := range rotationHints {
		sm.RotateToPosition(rh)
		if sm.Position == 0 {
			zeroCrossings++
		}
	}

	fmt.Printf("Zero crossings: %d\n", zeroCrossings)
	// ans: 997
}
