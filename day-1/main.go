package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

type Rotation struct {
	Direction Direction
	Steps     int
}

func parseRotation(line string) (Rotation, error) {
	if len(line) < 2 {
		return Rotation{}, fmt.Errorf("line too short: %s", line)
	}

	var dir Direction
	switch line[0] {
	case 'L':
		dir = Left
	case 'R':
		dir = Right
	default:
		return Rotation{}, fmt.Errorf("invalid direction: %c", line[0])
	}

	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid steps: %s", line[1:])
	}

	return Rotation{Direction: dir, Steps: steps}, nil
}

func parseFile(filename string) ([]Rotation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rotations []Rotation
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		rotation, err := parseRotation(line)
		if err != nil {
			fmt.Printf("Error parsing line '%s': %v\n", line, err)
			continue
		}

		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

func main() {
	rotations, err := parseFile("input.txt")

	if err != nil {
		fmt.Printf("Failed to parse file: %v\n", err)
		return
	}

	dial := 50

	password := 0
	password2 := 0

	fmt.Printf("Starting dial position: %d\n", dial)

	for i, rotation := range rotations {
		oldDial := dial

		zeroCrossings := 0

		// Find all 0 positions for each rotation instruction
		currentPos := dial
		for step := 0; step < rotation.Steps; step++ {
			switch rotation.Direction {
			case Left:
				currentPos--
				if currentPos < 0 {
					currentPos = 99
				}
			case Right:
				currentPos++
				if currentPos >= 100 {
					currentPos = 0
				}
			}

			if currentPos == 0 {
				zeroCrossings++
			}
		}

		password2 += zeroCrossings

		switch rotation.Direction {
		case Left:
			dial -= rotation.Steps
		case Right:
			dial += rotation.Steps
		}

		for dial < 0 {
			dial += 100
		}
		for dial >= 100 {
			dial -= 100
		}

		fmt.Printf("Step %d: %s%d -> %d to %d\n", i+1, rotation.Direction, rotation.Steps, oldDial, dial)

		if dial == 0 {
			password += 1
			fmt.Printf("  -> Landed on 0! Password increment: %d\n", password)
		}
	}

	fmt.Printf("Final dial position: %d\nPassword: %d\nPassword part 2: %d\n", dial, password, password2)
}
