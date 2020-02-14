package main

import (
	"strings"
	"strconv"
)

func Parser(lines []string) Input{
	var int_data []int

	info := strings.Split(lines[0], " ")
	data := strings.Split(lines[1], " ")

	max, _ := strconv.Atoi(info[0])
	number, _ := strconv.Atoi(info[1])

	for _, e := range data {
        n, err := strconv.Atoi(e)
        if err != nil {
            panic(err)
        }
        int_data = append(int_data, n)
    }

	input := Input{Max: max, Number: number, Slices: int_data}
	return input
} 
