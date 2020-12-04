package main

import "fmt"

type Position struct {
	X int
	Y int
}

func check_position (position Position, input []string, trees int) (int) {
	if position.Y == (len(input)-1) {
		return trees
	} else {
		position.X = position.X + 3
		if position.X > (len(input[0])-1) {
			position.X = position.X - (len(input[0]))
		}
		position.Y++
		if string(input[position.Y][position.X]) == "#" {
			trees++
		}
		return check_position(position, input, trees)
	}
	return 0
}

func part_one (input []string) {
	trees := 0
	position := Position{0,0}
	fmt.Println(check_position(position, input, trees))
}

func main () {
	given_input := []string {
"...............#...#..#...#....",
"...#....##.....##...######..#..",
"....#.....#.##..#...#..#.......",
"...#..........#.#....#......#..",
"....#.#...#.#.......#......##.#",
"....#.#..........#.....#.##....",
"##...#.#.##......#......#.#.#..",
"#.#.#........#....#..#.#.......",
"..#...##..#..#.......#....###..",
".#.#..........#...#........#..#",
".#..#......#....#.#...#...#.#.#",
"..#.........#..##.....#.#.##.#.",
".#......#...#....#.....#......#",
"........#..##..##.........#..#.",
".....#....###..#....##........#",
".###...#..##..#.##......##...##",
".#...#...#..#.#..#...##.....#..",
".......#....#....#.#...#.......",
".##.......#.##...#.#...#..#....",
"#.#...#......#....#.......#....",
"..###...............####...#.#.",
".##.#....#......#..#...#.#..#.#",
".............#.#.#......##.....",
"#....#.#.#........#....##...#..",
"...##....##...##..#...#........",
"..##......#...#......#...###...",
"...#...##......##.#.#..#.......",
"#......#..#...#.#..#......#..##",
".#..#..#........##....##.......",
".#...........###.###.##...#....",
"............#.#...........#.#..",
"#...#........#...#...#..#.#.#.#",
"...#.......#.#.#..#.#..........",
"......#..#..#....##..##..##....",
"........##...#......#..#.#.....",
"..#.#.......#......#...........",
"#.#.....#......##..........#.#.",
"#.....###.#...#...#..#....#...#",
".##.#...#............##.....#..",
"###....#.#.....#.......##......",
"##.......##.....#..#...#..##.##",
"....#.##............###...#..##",
".###..#...##.#..#..##..#.......",
".##.##..####.....#.........#...",
"....####..#...#....#.#...#.....",
"..##....#..#......#...........#",
"..........#......#..##.......#.",
".................#.#.#........#",
"#.......##.#...##.......##.##..",
".#................#.#.....##..#",
"......#..#............##.#.....",
"...##............#.....#.....#.",
"##.###.#.......#....#.........#",
"......#.....#.#.#.#......#.....",
"......#.##......#......##......",
"..#....#...#..#.....#..#....#.#",
".#.##.##.....#.......#.......#.",
"...#..#.#......##....##..#.....",
".#.....#......##...#..#.#....#.",
"..#......#....#..#..###.#.#....",
".....#........#.#...#..#.#.....",
"....#.#.......#...##.....####..",
"#..#..##...#...........#...#..#",
".#..#...#.....#.....#.#.....#.#",
"..##..###.....#...#.#.#.......#",
"#..##.##......###..#......###..",
"#..#...#.......#....#.#...#.#.#",
"...........###..#...#..##....#.",
".....#...........###.......#...",
"##......#.......#......##.#..#.",
"#.................#........#...",
"#.#.............#......#...#..#",
"......#.#....#....#....#....#.#",
"..#...#....#..#....#....#..#...",
"#....#......#..#...#..#....#.#.",
"..#.....#..#...#...#.......#...",
".#........###..##.#..#.........",
".....##.#.....#..###..#........",
"...#...#.###....######......#..",
".###.#..#.....#.....#..#...#...",
"##..#.#......#.........#...#..#",
"###...##..###.......#..##.#.#.#",
".#...................#..#...##.",
".#...#...###...#.......#......#",
"....#.....##....#...##.#...#.##",
"..#....#......#...#..###.......",
".........#..........##........#",
"...#.........##..#....#..###...",
"......#.##....#.#........#.#.##",
"....#..#...#.............#....#",
"#..#.....#.#.....#....#........",
"....#..#...####.#.###..#.......",
".....#....#....#..#..#.#.....#.",
"#..##....#.....#.#.............",
".##...#..#.......#...##.###....",
"##......#...##....#......##....",
"#......#.#......#.#..#......#..",
".#...#......#............###..#",
".#..#.#.....#.#.....#..#....#..",
".#............#...#..#..#...##.",
"...#...#....#.#.#....#....#....",
"........#......###.#....##.##.#",
"......#.#..................##..",
"..#..#........#..##..........##",
".#...#.#....#####.#.#..#.......",
".....#..#.#..#.....#.#........#",
"#.#..#..#.#..........#..#..#...",
".##........#.#.......#........#",
".....#....#..##...#......##....",
"....#.##.##...##.#.........#.#.",
"...#...#..#.#.......#.......#..",
".................##..#.........",
"..##.##....#...#.##.......#..#.",
"....#...........#.....###....##",
".#..................#..........",
"....#.##....#......##.#..#.##.#",
"#......#..#.#....##...####...#.",
"#.....#.#.##...........#.##...#",
".......#...##..#.........##....",
".#...#..........##......#..#.#.",
"#...#.#......#.....#........#..",
"........#.....#.......##......#",
".#..#...........#...#..#..#....",
"......#.##......##.#......#..##",
"......#....#..#................",
"##.#.......#......#..#....##.##",
"..#...###..#.......#...#.#....#",
".....#...#.........#.....#..#.#",
"..#.....#.........#..#...#.#...",
".#.........##..#.#...#.....##..",
"..........#.#.#...#....#....#..",
"....#.###.#..#..#..#.#........#",
"..#...##...##.#.#.....#.#..##..",
".#..#......#.####..#.......#..#",
"##.......#.....#.#.#..#...##..#",
".##......##...##.........#..#..",
"..#.##..#......#......#..#..#..",
"..#..#.##..#........#....#.#...",
"##.####...##.#..##.........#..#",
".#.......#.#..#.....#.##.......",
"........#.........#...........#",
"..#...#.####.....##.##.#....#.#",
".....#..##.#..###.##.#.#...#...",
"#..##..##....#...#....#...#....",
".###.#....#..#.......#......###",
".#....##.....#.#........#...#.#",
"#.#.#..#..#...#....#..#.....#..",
"....##...#.##..#............#..",
"##......###...#...#...#....#...",
"#.#...#.#..#..##.##..##..#....#",
".#...#.#....#..##.#..##..#.....",
".............#..#..#.#.....#...",
"#........#..........#....###...",
"..#..#......#...#........#.....",
".#..#.............#.........#..",
"#.....#....##..#..#.#.#..#.....",
"...#......#.........#.#....##.#",
"..#.......##....###..#.........",
".#.......#......#..............",
".#...##.....##...###.#....#.#..",
"......#....#.........#.......#.",
"##.......#..##....###.#.....##.",
".....#......####.....#......#..",
"....#....#..###....#.....##...#",
"...#...#.#........#.....#..#.##",
"#..#....#.#...###...##.....##..",
"#.....##...#.#............#....",
"##....#.......#.#.#....#.#.###.",
"#........#...#...####.#....#.#.",
".##.#......#........#..#.....#.",
"...##.#.......#...#.#.##..#...#",
"......#.........##..#....#.....",
".#......#...........#......#...",
"......#.#.#.#..#.#....#....#..#",
"##.................##...#.#....",
"........#.........#..#..#...###",
".#........#........#....#..#...",
"....###.##.##......#.#...#....#",
"#......#.#..............#......",
".......#..#....#..##.........#.",
"#............#....#............",
"#...#.#.........##..###...##...",
".#....#.#.#.....#..#..##.......",
".............##.#.#.......#..#.",
"#...#..##.#..###.....##..#.....",
"...#.#.....#...#......##.#..#..",
"#..........#.##.#...#...#.#...#",
"...#.#...#.........#.#..#.#..#.",
"#....#.................#.......",
"..#..#.#.#..#.#..##...#.#......",
"..#....#.#.#...#.....#...#.....",
"#...#.###......#.......#..#..#.",
"#.#.##.##..............#.#.#..#",
"#..........#.#.........#.###...",
"...........#.......#.#..#......",
"....#.#..#....#....#....##.....",
"#..#...##########.........#.#..",
"..#.............##........#.#..",
"#.#.##......#...#.#....##..##..",
"..##..#.#.#....#......#.#..#.#.",
".#.#...#................#......",
"#...#...#.....##.#...#.......#.",
".....##.......#...#......#.##..",
"....#.##..........#.#.##....##.",
"...........##........#.#.#.#...",
"..#...........###.#....#..#....",
"..#.#...#...#.#........#.....#.",
".##......##.....#........#.....",
"....#.......#....#...#.##...#..",
".#.....##.....#...##...........",
"#....#.##.##...#...###.#####...",
"..#.#......#.#.......#....#..#.",
"..........#...#...........#....",
".........#..#...#...#......#.##",
".#...#.#...#.###.##..#........#",
"#....#.....#.......##....#.....",
"#.....#..#.....#...##.......#..",
"#.#.#.#.............#.....#...#",
"...#.....#.....#...............",
"..##.###.#.#........#.........#",
"...##..#.##..#....#.#...#.#....",
"...##...#...##.#.........#.#..#",
".###......#....#.........#.#...",
"###.#.#...#.......#...#.....#.#",
"...#.#.......#.....####........",
"......#..#.....###.#....#..#...",
"..####...#...#..#......#...#...",
"#..............##....#......##.",
"..##..###...##..#.#.........#..",
"#..#.#...#.......#.........##..",
"..##....#......#...#..##.......",
"..#.#..#..###.....#.....#...###",
".#..#.##.....##.#.#.#........#.",
"..#.#.........#................",
"..#...........#.#...##.#...#..#",
".....#........#..#.....#....#..",
"#.#....#...........##.....#..##",
"##.......#.....#.....#.#......#",
".##............####.#........##",
"....##.#.....#......#....#...#.",
".#.#...##.#..##..#..........#..",
"..........................#....",
"##..##..#..#.#....#.....#......",
"...#.#........#.#.##.#.....#..#",
"#..#....#...#..#...#........#.#",
"#.......#####...#..#..#.#......",
"#.##....#......#......#..###...",
"..#.......#...........#.....##.",
"#...#....#..#......##...#......",
"...##.#..##........#..###......",
"##.#...........#.............##",
"#.....#..#.....#.....#.........",
".#..........#..#......###......",
"...#...#..##.......#...#...#.#.",
"..####......#.....#..#.#......#",
"....#..#.....#.#...............",
".#.......#....#.....#..##....#.",
".....#.........#..........##...",
"#...........#..#.....#........#",
"............#..##...#...#.#....",
"##.............####...#.....#..",
".....#......#....##.#.....##...",
"...........#.....#.#..#.#......",
".#.#......#....#.............##",
"##...#.#.......##........#.....",
"#..#.....#.#.......####...#..#.",
"....#.#...#....#.###..#..#.#...",
".....#.#..#......#.##.#####.#..",
".....#....##..........#......#.",
"#.......#........##.......##...",
"#...#..#..##.#....#...#...#....",
"...#..##..#...#..........#.....",
"#..#....###.....#......##......",
"...###......#.....#....#.......",
"#.#...#.#..###.####.......#.#.#",
"...#......#.#..##..#.....#.....",
"#.#............#.....##.#......",
"#..#......##...##..#...#.#..###",
"#.....#...#..#..#....#..###....",
"####..###....#..##....#..#.....",
"..##.#.....#.......##....#.#.#.",
"##..#.#.......#.#...##.#..#.#..",
"..#.#.#.##.#.#.#...........#...",
"...#.##.....#....#..#.#..#.....",
"...#..#.........#..........#..#",
"#...#..#.....#.#.#.............",
"##.#....##.##...#...#..#..#..#.",
"....####..##..##.#..#...##.....",
"##.....##.#.#...#.#.......###..",
"#..#.#....#......#.......##...#",
"#.#...............#..#..#......",
".....##...##..#........#......#",
".#..#............##......#....#",
"......#.#..#..##.#......#.....#",
"..#.#.............#...#......##",
"....#.#..#..#...##...#..##.....",
"#.#.............#...#.#..#....#",
"#..#..#.##....#....#...#.......",
"....#..#..........#.....##...#.",
"..#####.......#...#..........#.",
"....#......##.......#..##.#.#.#",
"#...#.#.....#....#....##.......",
"..##.#.#..#.#...#.....##.....#.",
"#.#..#....#............#..#.#..",
"...#.##..##..#.#...###......#.#",
"......##.......#....#......###.",
"....#..##.......#..#.#....#..#.",
"#............#..........##..#.#",
"..#.....#..#.#..##..#....##.#..",
".....#.....#....##.#.#......#.#",
"...####.......###..#...#.#....#",
".#.##.....##.....##..#..#......",
"...#..###..##..................",
"##..##.....#.#..#..#.#........#",
".#.#........##.........#....#..",
"........#......###....#...#....",
"......#...........#...........#",
".#...###...#.........#.#...#..#",
".....#..........#......##....#.",
".#.#...#........##.......#.#...",
".....##.....##....#...#...#..#.",
"..#.....................##...##",
"#..#.#......##...##..#......#.."}

part_one(given_input)
}

