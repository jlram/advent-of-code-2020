// desc: https://adventofcode.com/2020/day/6

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func partOne(input []string) int {
	count := 0

	for _, group := range input {
		m := make(map[rune]bool)
		for _, letter := range strings.Replace(strings.Replace(group, "\r", "", -1), "\n", "", -1) {
			if !m[letter] {
				m[letter] = true
			}
		}
		count += len(m)
	}
	return count
}

func partTwo(input []string) int {
	count := 0

	for _, group := range input {
		m := make(map[rune]int)

		for _, letter := range strings.Replace(strings.Replace(group, "\r", "", -1), "\n", "", -1) {
			m[letter]++
		}

		for _, value := range m {
			if value == len(strings.Split(group, "\r\n")) {
				count++
			}
		}
	}
	return count
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n\r")

	fmt.Println(partOne(givenInput))
	fmt.Println(partTwo(givenInput))
}
