package main

import (
	"fmt"
	"os"

	common "github.com/halamix2/advent_of_code/cmd/02"
)

func main() {
	// idRanges, err := common.LoadInput("cmd/02/in.txt")
	idRanges, err := common.LoadInput("cmd/02/input.txt")
	if err != nil {
		fmt.Printf("Couldn't load IDs: %v\n", err)
		os.Exit(1)
	}
	sumInvalid := 0
	for idRange := range idRanges {
		invalidIDs := idRange.GetAllInvalid(0)
		for _, invalid := range invalidIDs {
			sumInvalid += invalid
		}
	}
	fmt.Printf("Sum of invalid IDs: %d\n", sumInvalid)
}
