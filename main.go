package main

import (
	"sort"
)

func main() {
	files := []string{"a_example", "b_small", "c_medium", "d_quite_big", "e_also_big"}

	for _, file := range files {
		lines := Read(file)
		pizza := Parser(lines)
		output := calculatePizza(pizza)
		Write(file, output)
	}
}

func calculatePizza(input Input) Output{
	var output Output

	total := 0
	for index := input.Number - 1; index >= 0; index-- {
		slice := input.Slices[index]
		if (slice + total) <= input.Max {
			output.Indexes = append(output.Indexes, index)
			total += slice
		}
	} 

	sort.Ints(output.Indexes)
	output.Number = len(output.Indexes)

	return output
}