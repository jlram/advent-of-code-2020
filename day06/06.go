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
		m := make(map[string]int)
		for _, letter := range strings.Replace(strings.Replace(group, "\r", "", -1), "\n", "", -1) {
			if _, ok := m[string(letter)]; !ok {
				m[string(letter)] = 0
			}
		}
		count += len(m)
	}
	return count
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n\r")

	fmt.Println(partOne(givenInput))
}
