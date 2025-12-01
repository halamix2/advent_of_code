package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type RotationDirection int

const (
	RotateLeft RotationDirection = iota
	RotateRight
)

type RotationHint struct {
	direction RotationDirection
	step      int
}

func (rh RotationHint) String() string {
	return fmt.Sprintf("dir: %b; step: %d", rh.direction, rh.step)
}

func LoadInput(filePath string) ([]RotationHint, error) {
	rotateHints := make([]RotationHint, 0)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %s\n", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rot := scanner.Text()[0]
		if rot != 'R' && rot != 'L' {
			return nil, fmt.Errorf("unknown rotation direction: %b", rot)
		}
		step, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			return nil, fmt.Errorf("can't convert step to numberr: %w", err)
		}
		rh := RotationHint{step: step}
		if rot == 'L' {
			rh.direction = RotateLeft
		} else {
			rh.direction = RotateRight
		}
		rotateHints = append(rotateHints, rh)

	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %s", err)
	}

	return rotateHints, f.Close()
}

type SafeMachine struct {
	Position int
}

func (sm *SafeMachine) RotateToPosition(rh RotationHint) {
	newPosition := sm.Position
	// sometimes we got maaany rotations, so we simplify here
	actualStep := rh.step % 100
	if rh.direction == RotateLeft {
		newPosition -= actualStep
		for newPosition < 0 {
			newPosition += 100
		}
	} else {
		newPosition += actualStep
		if newPosition > 99 {
			newPosition -= 100
		}
	}
	sm.Position = newPosition
}

func (sm *SafeMachine) ClickToPosition(rh RotationHint, clickPosition int) int {
	// slow bruteforce, yummy
	clicks := 0
	position := sm.Position
	steps := rh.step
	for range steps {

		if rh.direction == RotateLeft {
			position--
		} else {
			position++
		}
		if position == 100 {
			position = 0
		}
		if position == -1 {
			position = 99
		}
		if position == clickPosition {
			clicks += 1
		}
	}

	sm.Position = position
	return clicks
}
