package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		var nums []int64

		line := scanner.Text()

		for {
			earliestIndex := len(line)
			earliestKey := ""

			for k, _ := range numberMap {
				index := strings.Index(line, k)
				if index != -1 && index < earliestIndex {
					earliestIndex = index
					earliestKey = k
				}
			}

			if earliestKey == "" {
				break
			}

			line = strings.Replace(line, earliestKey, numberMap[earliestKey], 1)
		}

		chars := strings.Split(line, "")
		for _, char := range chars {
			num, err := strconv.ParseInt(char, 0, 10)
			if err == nil {
				nums = append(nums, num)
			}
		}

		val, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[0]))
		if err == nil {
			sum += val
		}
	}

	// total
	fmt.Print(sum)
}
