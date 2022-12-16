package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day-4/assignment-pairs.txt")
	defer file.Close()

	overlapCountPart1, overlapCountPart2 := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		firstElf, secondElf := splitLine[0], splitLine[1]
		firstElfStart, firstElfEnd := strings.Split(firstElf, "-")[0], strings.Split(firstElf, "-")[1]
		secondElfStart, secondElfEnd := strings.Split(secondElf, "-")[0], strings.Split(secondElf, "-")[1]

		firstElfStartInt, _ := strconv.ParseInt(firstElfStart, 10, 0)
		firstElfEndInt, _ := strconv.ParseInt(firstElfEnd, 10, 0)
		secondElfStartInt, _ := strconv.ParseInt(secondElfStart, 10, 0)
		secondElfEndInt, _ := strconv.ParseInt(secondElfEnd, 10, 0)

		if (secondElfStartInt >= firstElfStartInt && secondElfEndInt <= firstElfEndInt) ||
			(firstElfStartInt >= secondElfStartInt && firstElfEndInt <= secondElfEndInt) {
			overlapCountPart1 = overlapCountPart1 + 1
		}

		// (firstS <= secondS and firstE => secondE)
		// (firstS => secondS and firstE <= secondE)
		if (secondElfStartInt >= firstElfStartInt && secondElfEndInt <= firstElfEndInt) ||
			(firstElfStartInt >= secondElfStartInt && firstElfEndInt <= secondElfEndInt) ||
			(firstElfStartInt == secondElfEndInt || firstElfEndInt == secondElfStartInt) {
			fmt.Print("overlap 2: ", firstElf, " --- ", secondElf, "\n")
			overlapCountPart2 = overlapCountPart2 + 1
		}

	}
	fmt.Print("total overlap count partido uno: ", overlapCountPart1, "\n")
	fmt.Print("total overlap count partido dos: ", overlapCountPart2, "\n")
}
