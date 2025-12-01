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

	zeroClicks := 0
	clickPosition := 0
	sm := common.SafeMachine{Position: 50}
	for _, rh := range rotationHints {
		zeroClicks += sm.ClickToPosition(rh, clickPosition)
	}

	fmt.Printf("Zero clicks: %d\n", zeroClicks)
}
