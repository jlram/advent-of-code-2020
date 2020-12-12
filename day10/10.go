package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(input map[int]bool) int {
	value := 0
	oneCounter := 0
	threeCounter := 0
	loop := true

	for loop {
		if input[value+1] {
			value++
			oneCounter++
		} else if input[value+2] {
			value = value + 2
		} else if input[value+3] {
			value = value + 3
			threeCounter++
		} else {
			loop = false
		}
	}

	fmt.Println(oneCounter, threeCounter+1)

	return oneCounter * (threeCounter + 1)
}

func partTwo(input map[int]bool) int {
	return 0
}

func main() {
	inputMap := make(map[int]bool)
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		inputMap[num] = true
	}
	fmt.Println(partOne(inputMap))
	fmt.Println(partTwo(inputMap))
}
