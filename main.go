package main

import (
	"fmt"
)

func main() {
	day1 := fetch("https://adventofcode.com/2021/day/1/input")
	day2 := fetch("https://adventofcode.com/2021/day/2/input")

	fmt.Println("Day 1 - Part 1:", dayOnePartOne(day1))
	fmt.Println("Day 1 - Part 2:", dayOnePartTwo(day1))

	fmt.Println("Day 2 - Part 1:", dayTwoPartOne(day2))
	fmt.Println("Day 2 - Part 2:", dayTwoPartTwo(day2))
}
