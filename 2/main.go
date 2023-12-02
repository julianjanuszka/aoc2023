package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxMarblesMap = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	powerSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		colonIndex := strings.Index(line, ":")
		gameString := line[0:colonIndex]
		gameDigit := strings.Trim(gameString, "Game ")
		gameInt, err := strconv.Atoi(gameDigit)
		if err != nil {
			continue
		}

		lineValid := true
		var minMarbles = make(map[string]int)
		for {
			nextSemiColonIdx := strings.Index(line, ";")
			var currentMarbleString string
			if nextSemiColonIdx == -1 {
				currentMarbleString = line
			} else {
				currentMarbleString = line[0:nextSemiColonIdx]
			}

			valid := isMarblePullValid(currentMarbleString, minMarbles)
			if !valid {
				lineValid = false
			}

			if nextSemiColonIdx == -1 {
				// calculate power (pt2)
				power := 0
				for _, v := range minMarbles {
					if power == 0 {
						power = v
					} else {
						power = power * v
					}
				}
				powerSum += power
				break
			}
			line = line[nextSemiColonIdx+1:]
		}
		if lineValid {
			sum += gameInt
		}
	}
	fmt.Printf("sum %d, power sum %d", sum, powerSum)
}

func isMarblePullValid(pull string, minMap map[string]int) bool {
	isValid := true
	options := []string{"red", "green", "blue"}
	for _, color := range options {
		if strings.Contains(pull, color) {
			index := strings.Index(pull, color)
			stringDigit := strings.Trim(pull[index-3:index-1], " ")
			digit, err := strconv.Atoi(stringDigit)
			if err == nil {
				currentMin := minMap[color]
				if digit > currentMin {
					minMap[color] = digit
				}

				maxMarbles := maxMarblesMap[color]
				if digit > maxMarbles {
					isValid = false
				}
			}
		}
	}
	return isValid
}
