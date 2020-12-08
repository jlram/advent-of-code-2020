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
			} else if number, _ := strconv.Atoi(strings.ReplaceAll(value[1], "\r", "")); value[0] == "acc" {
				acc += number
				index++
			} else if value[0] == "jmp" {
				index += number
			}
		}
	}
	return acc
}

func partTwo(input []string) int {
	var possibleArr []int
	acc := 0

	for i, v := range input {
		if strings.Contains(v, "nop") || strings.Contains(v, "jmp") {
			possibleArr = append(possibleArr, i)
		}
	}

	for _, v := range possibleArr {
		acc = 0
		index := 0
		exit := false
		aux := make([]string, len(input))
		copy(aux, input)
		var arr []int

		// Changes "nop" to "jmp" and viceversa
		if strings.Split(aux[v], " ")[0] == "nop" {
			aux[v] = "jmp " + strings.Split(aux[v], " ")[1]
		} else if strings.Split(aux[v], " ")[0] == "jmp" {
			aux[v] = "nop " + strings.Split(aux[v], " ")[1]
		}

		for !exit {
			if index == len(aux)-1 { // If reaches the end of the input, returns acc value
				return acc
			}

			if find(arr, index) {
				exit = true // iterates next value
			} else {
				arr = append(arr, index)
				value := strings.Split(aux[index], " ")
				if value[0] == "nop" {
					index++
				} else if number, _ := strconv.Atoi(strings.ReplaceAll(value[1], "\r", "")); value[0] == "acc" {
					acc += number
					index++
				} else if value[0] == "jmp" {
					index += number
				}
			}
		}
	}
	return 0
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\r\n")

	fmt.Println(partOne(givenInput))
	fmt.Println(partTwo(givenInput))
}
