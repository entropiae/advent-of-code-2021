package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readMeasurements(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var measurements []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()

		m, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Cannot convert value: %s", err)
		}
		measurements = append(measurements, m)
	}
	return measurements, scanner.Err()
}

func sumArray(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func main() {
	fname := os.Args[1]
	measurements, err := readMeasurements(fname)

	if err != nil {
		os.Exit(-1)
	}

	increases := -1
	prevMeasurement := 0

	for _, m := range measurements {
		//fmt.Printf("Current: %d, previous: %d\n", m, prevMeasurement)
		if m > prevMeasurement {
			increases++
			//fmt.Printf("Increased! %d\n", increases)
		}
		prevMeasurement = m
	}
	fmt.Printf("Part One: increased measurements: %d\n", increases)

	increases = -1
	prevMeasurement = 0

	for i := 0; i <= len(measurements)-3; i++ {
		m := sumArray(measurements[i : i+3])

		//fmt.Printf("Current: %d, previous: %d\n", m, prevMeasurement)
		if m > prevMeasurement {
			increases++
			//fmt.Printf("Increased! %d\n", increases)
		}
		prevMeasurement = m
	}
	fmt.Printf("Part Two: increased measurements: %d\n", increases)

}
