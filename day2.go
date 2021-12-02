package main

func dayTwoPartOne(directions []string) int {
	pos := 0
	depth := 0

	for i := 0; i < len(directions); i += 2 {
		direction := strToInt(directions[i+1])

		switch directions[i] {
		case "forward":
			pos += direction
		case "down":
			depth += direction
		case "up":
			depth -= direction
		}
	}

	return pos * depth
}

func dayTwoPartTwo(directions []string) int {
	aim := 0
	pos := 0
	depth := 0

	for i := 0; i < len(directions); i += 2 {
		direction := strToInt(directions[i+1])

		switch directions[i] {
		case "forward":
			pos += direction
			depth += direction * aim
		case "down":
			aim += direction
		case "up":
			aim -= direction
		}
	}

	return pos * depth
}
