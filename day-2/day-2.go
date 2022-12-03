package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

// A, X == Rock 1, B, Y == Paper 2, C, Z == Scissors 3
// win == 6, draw == 3, lose == 0

func main() {
    file, _ := os.Open("day-2/strategy.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        split_line := strings.Split(line, " ")
        opponent_move, my_move := split_line[0], split_line[1]
    }
}