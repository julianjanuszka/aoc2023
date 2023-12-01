package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		chars := strings.Split(line, "")
		for _, char := range chars {
			num, err := strconv.ParseInt(char, 0, 10)
			if err == nil {
				nums = append(nums, num)
			}
		}

		var val int
		if len(nums) == 1 {
			val, err = strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[0]))
		} else {
			val, err = strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
		}
		if err == nil {
			sum += val
		}
	}

	// total
	fmt.Print(sum)
}
