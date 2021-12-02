package main

func dayOnePartOne(input []string) int {
	depths := strArrToIntArr(input)
	prev := depths[0]
	inc := 0

	for _, depth := range depths {
		if depth > prev {
			inc++
		}
		prev = depth
	}

	return inc
}

func dayOnePartTwo(input []string) int {
	depths := strArrToIntArr(input)
	prev := sum(depths[:3])
	inc := 0

	for i := 0; i < len(depths); i++ {
		window := sum(depths[i : i+3])
		if window > prev {
			inc++
		}
		prev = window
	}

	return inc
}
