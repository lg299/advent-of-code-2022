package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strconv"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func main() {
    file, err := os.Open("day-1/elf-calories.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var total_elf_calories []int
    counter := 1
    elf_calories := 0
    for scanner.Scan() {
    	line := scanner.Text()
        calories, _ := strconv.Atoi(line)
        elf_calories = elf_calories + calories

    	if len(line) <= 0 || counter == 2223 {
    	    total_elf_calories = append(total_elf_calories, elf_calories,)
    	    elf_calories = 0
    	}

        counter++

    }

    _, max := MinMax(total_elf_calories)
    fmt.Print(max)
}

