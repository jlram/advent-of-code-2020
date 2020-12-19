package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Ship struct {
	Direction rune
	NSBalance int
	EWBalance int
}

type Waypoint struct {
	NSBalance int
	EWBalance int
}

func partOne(input []string) int {
	ship := Ship{'E', 0, 0}
	for _, v := range input {
		numValue, _ := strconv.Atoi(v[1:])
		switch v[0] {
		case 'N':
			ship.NSBalance += numValue
		case 'S':
			ship.NSBalance -= numValue
		case 'E':
			ship.EWBalance += numValue
		case 'W':
			ship.EWBalance -= numValue
		case 'R', 'L':
			ship.Direction = rotateElement(ship.Direction, rune(v[0]), numValue)
		case 'F':
			switch ship.Direction {
			case 'N':
				ship.NSBalance += numValue
			case 'S':
				ship.NSBalance -= numValue
			case 'E':
				ship.EWBalance += numValue
			case 'W':
				ship.EWBalance -= numValue
			}
		}
	}
	return ManhattanDistance(ship.NSBalance, ship.EWBalance)
}

func partTwo(input []string) int {
	ship := Ship{'E', 0, 0}
	waypoint := Waypoint{1, 10}
	for _, v := range input {
		numValue, _ := strconv.Atoi(v[1:])
		switch v[0] {
		case 'N':
			waypoint.NSBalance += numValue
		case 'S':
			waypoint.NSBalance -= numValue
		case 'E':
			waypoint.EWBalance += numValue
		case 'W':
			waypoint.EWBalance -= numValue
		case 'R', 'L':
			waypoint = rotateWaypoint(waypoint, ship.Direction, rune(v[0]), numValue)
		case 'F':
			ship.NSBalance += waypoint.NSBalance * numValue
			ship.EWBalance += waypoint.EWBalance * numValue
		}
	}

	return ManhattanDistance(ship.NSBalance, ship.EWBalance)
}

func ManhattanDistance(x int, y int) int {
	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

func rotateElement(direction rune, course rune, degrees int) rune {
	rotate90Right := map[rune]rune{'N': 'E', 'S': 'W', 'E': 'S', 'W': 'N'}
	rotate270Right := map[rune]rune{'N': 'W', 'S': 'E', 'E': 'N', 'W': 'S'}
	rotate180 := map[rune]rune{'N': 'S', 'S': 'N', 'E': 'W', 'W': 'E'}

	if (course == 'R' && degrees == 90) || (course == 'L' && degrees == 270) {
		return rotate90Right[direction]
	} else if (course == 'R' && degrees == 270) || (course == 'L' && degrees == 90) {
		return rotate270Right[direction]
	} else {
		return rotate180[direction]
	}
}

func rotateWaypoint(waypoint Waypoint, direction rune, course rune, degrees int) Waypoint {
	if (course == 'R' && degrees == 90) || (course == 'L' && degrees == 270) {
		aux := waypoint.EWBalance
		waypoint.EWBalance = waypoint.NSBalance
		waypoint.NSBalance = -aux
	} else if (course == 'R' && degrees == 270) || (course == 'L' && degrees == 90) {
		aux := waypoint.NSBalance
		waypoint.NSBalance = waypoint.EWBalance
		waypoint.EWBalance = -aux
	} else if degrees == 180 {
		waypoint.NSBalance = -waypoint.NSBalance
		waypoint.EWBalance = -waypoint.EWBalance
	}

	return waypoint
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n")

	fmt.Println(partOne(givenInput))
	fmt.Println(partTwo(givenInput))
}
