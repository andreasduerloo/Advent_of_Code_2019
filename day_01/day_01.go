package day_01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() (int, int) {
	file, err := os.Open("./inputs/01.txt")

	if err != nil {
		panic("Something went wrong with the file - bye!")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	weights := make([]int, 0)

	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Not having a good time converting" + scanner.Text() + "to an int.")
			continue
		}

		weights = append(weights, weight)
	}

	first := mapReduce(weights, fuelFunc, func(a, b int) int { return a + b }, 0)
	second := mapReduce(weights, totalFuel, func(a, b int) int { return a + b }, 0)

	return first, second
}

func fuelFunc(weight int) int {
	return (weight / 3) - 2
}

func totalFuel(weight int) int {
	return recFuel(weight, 0)
}

func recFuel(weight, fuel int) int {
	if fuelFunc(weight) <= 0 {
		return fuel
	} else {
		return recFuel(fuelFunc(weight), fuel+fuelFunc(weight))
	}
}

func mapReduce(ints []int, mapfunc func(int) int, reducefunc func(int, int) int, acc int) int {
	for _, elem := range ints {
		acc = reducefunc(acc, mapfunc(elem))
	}

	return acc
}
