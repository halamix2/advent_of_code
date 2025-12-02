package common

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	firstID int
	lastID  int
}

func (ir IDRange) String() string {
	return fmt.Sprintf("f: %d, l: %d", ir.firstID, ir.lastID)
}

func (ir *IDRange) GetAllInvalid(staticRepeats int) []int {
	allInvalid := make([]int, 0)
	for i := ir.firstID; i <= ir.lastID; i++ {
		idString := strconv.Itoa(i)
		fullLength := len(idString)
		halfLength := fullLength / 2
		// now: up to half length repeat first n characters to match the full length
		for n := 1; n <= halfLength; n++ {
			// repeats is how many repeats can we push in, static 2 for the first task
			repeats := staticRepeats
			if repeats == 0 {
				repeats = fullLength / n
			}
			if n*repeats != fullLength {
				// this is 100% valid ID
				continue
			}
			prefix := idString[0:n]
			if strings.Repeat(prefix, repeats) == idString {
				allInvalid = append(allInvalid, i)
				break
			}
		}
	}
	return allInvalid
}

func LoadInput(filePath string) (iter.Seq[IDRange], error) {
	idRanges := make([]IDRange, 0)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %s\n", err)
	}
	ranges := strings.Split(string(fileData[:len(fileData)-1]), ",")
	for _, r := range ranges {
		IDs := strings.Split(r, "-")
		firstID, err := strconv.Atoi(IDs[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert first ID: %s\n", err)
		}
		lastID, err := strconv.Atoi(IDs[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert last ID: %s\n", err)
		}

		idRange := IDRange{firstID: firstID, lastID: lastID}
		idRanges = append(idRanges, idRange)
	}

	return func(yield func(IDRange) bool) {
		for _, idRange := range idRanges {
			if !yield(idRange) {
				return
			}
		}
	}, nil
}
