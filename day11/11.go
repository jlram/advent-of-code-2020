package main

import (
	"bufio"
	"fmt"
	"os"
)

func partOne(input [][]rune) int {
	changed := true
	counter := 0
	totalOccupied := 0

	for changed {
		counter++
		changed = false
		for rowKey, row := range input {
			for seatKey := range row {
				if rowKey == 0 && seatKey == 0 { // top left corner
					if free := 0; input[rowKey][seatKey] == 'L' {
						if input[rowKey+1][seatKey+1] != '#' {
							free++
						}
						if input[rowKey][seatKey+1] != '#' {
							free++
						}
						if input[rowKey+1][seatKey] != '#' {
							free++
						}
						if free == 3 {
							input[rowKey][seatKey] = '#'
							totalOccupied++
							changed = true
						}
					} else if occupied := 0; input[rowKey][seatKey] == '#' {
						if input[rowKey+1][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey+1][seatKey] == '#' {
							occupied++
						}
						if occupied >= 4 {
							input[rowKey][seatKey] = 'L'
							totalOccupied--
							changed = true
						}
					}
				} else if rowKey == 0 && seatKey == len(row)-1 { // top right corner
					if free := 0; input[rowKey][seatKey] == 'L' {
						if input[rowKey+1][seatKey-1] != '#' {
							free++
						}
						if input[rowKey][seatKey-1] != '#' {
							free++
						}
						if input[rowKey+1][seatKey] != '#' {
							free++
						}
						if free == 3 {
							input[rowKey][seatKey] = '#'
							totalOccupied++
							changed = true
						}
					} else if occupied := 0; input[rowKey][seatKey] == '#' {
						if input[rowKey+1][seatKey-1] == '#' {
							occupied++
						}
						if input[rowKey][seatKey-1] == '#' {
							occupied++
						}
						if input[rowKey+1][seatKey] == '#' {
							occupied++
						}
						if occupied >= 4 {
							input[rowKey][seatKey] = 'L'
							totalOccupied--
							changed = true
						}
					}
				} else if rowKey == len(input)-1 && seatKey == 0 { // bottom left corner
					if free := 0; input[rowKey][seatKey] == 'L' {
						if input[rowKey-1][seatKey+1] != '#' {
							free++
						}
						if input[rowKey][seatKey+1] != '#' {
							free++
						}
						if input[rowKey-1][seatKey] != '#' {
							free++
						}
						if free == 3 {
							input[rowKey][seatKey] = '#'
							totalOccupied++
							changed = true
						}
					} else if occupied := 0; input[rowKey][seatKey] == '#' {
						if input[rowKey-1][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey-1][seatKey] == '#' {
							occupied++
						}
						if occupied >= 4 {
							input[rowKey][seatKey] = 'L'
							totalOccupied--
							changed = true
						}
					}
				} else if rowKey == len(input)-1 && seatKey == len(row)-1 { // bottom right corner
					if free := 0; input[rowKey][seatKey] == 'L' {
						if input[rowKey-1][seatKey-1] != '#' {
							free++
						}
						if input[rowKey][seatKey-1] != '#' {
							free++
						}
						if input[rowKey-1][seatKey] != '#' {
							free++
						}
						if free == 3 {
							input[rowKey][seatKey] = '#'
							totalOccupied++
							changed = true
						}
					} else if occupied := 0; input[rowKey][seatKey] == '#' {
						if input[rowKey-1][seatKey-1] == '#' {
							free++
						}
						if input[rowKey][seatKey-1] == '#' {
							free++
						}
						if input[rowKey-1][seatKey] == '#' {
							free++
						}
						if occupied >= 4 {
							input[rowKey][seatKey] = 'L'
							totalOccupied--
							changed = true
						}
					}
				} else if rowKey == 0 { // first row
					if free := 0; input[rowKey][seatKey] == 'L' {
						if input[rowKey][seatKey+1] != '#' {
							free++
						}
						if input[rowKey][seatKey-1] != '#' {
							free++
						}
						if input[rowKey+1][seatKey+1] != '#' {
							free++
						}
						if input[rowKey+1][seatKey] != '#' {
							free++
						}
						if input[rowKey+1][seatKey-1] != '#' {
							free++
						}
						if free == 5 {
							input[rowKey][seatKey] = '#'
							totalOccupied++
							changed = true
						}
					} else if occupied := 0; input[rowKey][seatKey] == '#' {
						if input[rowKey][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey][seatKey-1] == '#' {
							occupied++
						}
						if input[rowKey+1][seatKey+1] == '#' {
							occupied++
						}
						if input[rowKey+1][seatKey] == '#' {
							occupied++
						}
						if input[rowKey+1][seatKey-1] == '#' {
							occupied++
						}
						if occupied >= 4 {
							input[rowKey][seatKey] = 'L'
							totalOccupied--
							changed = true
						}
					}
				} else if free, occupied := 0, 0; seatKey == 0 { // first seat in row
					if input[rowKey][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if free == 5 && input[rowKey][seatKey] == 'L' {
						input[rowKey][seatKey] = '#'
						totalOccupied++
						changed = true
					} else if occupied >= 4 && input[rowKey][seatKey] == '#' {
						input[rowKey][seatKey] = 'L'
						totalOccupied--
						changed = true
					}
				} else if free, occupied := 0, 0; rowKey == len(input)-1 { // last row
					if input[rowKey][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if free == 5 && input[rowKey][seatKey] == 'L' {
						input[rowKey][seatKey] = '#'
						totalOccupied++
						changed = true
					} else if occupied >= 4 && input[rowKey][seatKey] == '#' {
						input[rowKey][seatKey] = 'L'
						totalOccupied--
					}
				} else if free, occupied := 0, 0; seatKey == len(row)-1 { // last seat in row
					if input[rowKey][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if free == 5 && input[rowKey][seatKey] == 'L' {
						input[rowKey][seatKey] = '#'
						totalOccupied++
						changed = true
					} else if occupied >= 4 && input[rowKey][seatKey] == '#' {
						input[rowKey][seatKey] = 'L'
						totalOccupied--
						changed = true
					}
				} else { // regular seat
					free, occupied := 0, 0
					if input[rowKey][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey-1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey+1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey+1][seatKey-1] != '#' {
						free++
					} else {
						occupied++
					}

					if input[rowKey][seatKey] == '#' {
						fmt.Println(occupied, free, counter, rowKey, seatKey)
					}

					if free == 8 && input[rowKey][seatKey] == 'L' {
						input[rowKey][seatKey] = '#'
						totalOccupied++
						changed = true
					} else if occupied >= 4 && input[rowKey][seatKey] == '#' {
						input[rowKey][seatKey] = 'L'
						totalOccupied--
						changed = true
					}
				}
			}
		}
		fmt.Println(changed)
	}

	return totalOccupied
}

func main() {
	var input [][]rune
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		var arr []rune
		text := s.Text()
		for _, char := range text {
			arr = append(arr, char)
		}
		input = append(input, arr)
	}

	fmt.Println(partOne(input))
}
