// desc: https://adventofcode.com/2020/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func partOne(input []string) int {
	acc := 0
	index := 0
	var arr []int
	exit := false

	for !exit {
		if find(arr, index) {
			exit = true
		} else {
			arr = append(arr, index)
			value := strings.Split(input[index], " ")
			if value[0] == "nop" {
				index++
			} else if number, _ := strconv.Atoi(strings.ReplaceAll(value[1][1:], "\r", "")); value[0] == "acc" {
				if value[1][0] == '+' { // sum symbol
					acc += number
					index++
				} else { // subtract
					acc = acc - number
					index++
				}
			} else if value[0] == "jmp" {
				if value[1][0] == '+' { // sum symbol
					index += number
				} else { // subtract
					index = index - number
				}
			}
		}
	}
	return acc
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n")

	fmt.Println(partOne(givenInput))
}
