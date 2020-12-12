package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(input []int) int {
	min, max := 0, 25
	toIterate := input[25:]

	for _, value := range toIterate {
		arr := make([]int, 25)
		copy(arr, input[min:max])
		valid := false
		for _, number := range arr {
			for _, number2 := range arr {
				if number+number2 == value && number != number2 {
					valid = true
				}
			}
		}

		min++
		max++

		if !valid {
			return value
		}
	}
	return 0
}

func calculateTwo(input []int, value int) []int {
	count := 1
	min, max := 0, count
	valid := false
	arr := make([]int, count)
	for !valid {
		arr = make([]int, count)
		copy(arr, input[min:max+1])
		for _, number := range arr {
			for _, number2 := range arr {
				if number+number2 == value && number != 0 && number2 != 0 {
					valid = true
					fmt.Println(number, number2)
					return arr
				}
			}
		}
		min++
		max++
		if max == len(input)-1 {
			count++
			min, max = 0, count
		}
	}
	return arr
}

// TODO panic: runtime error: slice bounds out of range [:1025] with capacity 1024
func partTwo(input []int, value int) int {

	arr := calculateTwo(input, value)

	minValue, maxValue := 2147483647, 0
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}

		if v < minValue && v != 0 {
			minValue = v
		}
	}

	fmt.Println(minValue, maxValue, len(arr))
	return maxValue + minValue
}

func main() {
	var intInput []int
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		intInput = append(intInput, num)
	}
	resultOne := partOne(intInput)
	fmt.Println(resultOne)
	fmt.Println(partTwo(intInput, resultOne))
}
