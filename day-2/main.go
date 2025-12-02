package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func parseRange(inputString string) (*Range, error) {
	parts := strings.Split(inputString, "-")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid range format: %s", inputString)
	}

	start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("invalid start of range: %s", parts[0])
	}

	end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return nil, fmt.Errorf("invalid end of range: %s", parts[1])
	}

	return &Range{Start: start, End: end}, nil
}

func parseFile(filename string) ([]Range, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ranges []Range
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		inputRanges := strings.Split(line, ",")
		for _, inputRange := range inputRanges {
			parsedRange, err := parseRange(inputRange)
			if err != nil {
				fmt.Printf("Error parsing range '%s': %v\n", inputRange, err)
				continue
			}
			ranges = append(ranges, *parsedRange)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ranges, nil
}

func consistsOfRepeatedSequences(numberString string) bool {
	length := len(numberString)

	if length == 1 {
		return false
	}

	for seqLen := 1; seqLen <= length/2; seqLen++ {
		if length%seqLen == 0 {
			sequence := numberString[0:seqLen]
			isValid := true

			for pos := seqLen; pos < length; pos += seqLen {
				if numberString[pos:pos+seqLen] != sequence {
					isValid = false
					break
				}
			}

			if isValid && length/seqLen >= 2 {
				return true
			}
		}
	}

	return false
}

func main() {
	ranges, err := parseFile("input.txt")

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sum1 := 0
	sum2 := 0

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			numberString := strconv.Itoa(i)
			digits := len(numberString)

			// Part 1: Check if first half equals second half (even digits only)
			if digits%2 == 0 {
				firstSequence, err := strconv.Atoi(strings.TrimSpace(numberString[:digits/2]))
				if err != nil {
					fmt.Printf("Error parsing first half of %s: %v\n", numberString, err)
					continue
				}

				secondSequence, err := strconv.Atoi(strings.TrimSpace(numberString[digits/2:]))
				if err != nil {
					fmt.Printf("Error parsing second half of %s: %v\n", numberString, err)
					continue
				}

				if firstSequence == secondSequence {
					sum1 += i
				}
			}

			// Part 2: Check for repeated sequences
			if consistsOfRepeatedSequences(numberString) {
				sum2 += i
			}
		}
	}

	fmt.Printf("Sum of all invalid IDs (Part 1): %d\nSum of all invalid IDs (Part 2): %d\n", sum1, sum2)
}
