package main

import (
	"bufio"
	"os"
	"strconv"

)

func Read(filename string) []string {
	file, err := os.Open("./input/" + filename + ".in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Write(filename string, output Output) {
	file, err := os.Create("./output/" + filename + ".out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	s := strconv.Itoa(output.Number) 
	writer.WriteString(s + "\n")

	for _, index := range output.Indexes {
		i := strconv.Itoa(index) 
		writer.WriteString(i + " ")
	}

	writer.Flush()
} 