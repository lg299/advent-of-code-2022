package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A, X == Rock 1, B, Y == Paper 2, C, Z == Scissors 3
// win == 6, draw == 3, lose == 0

func main() {
	totalPointsScored := 0
	letterConversion := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	part2LetterConversion := map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	movePoints := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	file, _ := os.Open("day-2/strategy.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pointsScored := 0
		splitLine := strings.Split(line, " ")

		// part 2
		opponentMove, outcome := letterConversion[splitLine[0]], part2LetterConversion[splitLine[1]]
		if outcome == "draw" {
			pointsScored = 3 + movePoints[opponentMove]
		} else if outcome == "win" {
			if opponentMove == "rock" {
				pointsScored = 6 + 2
			} else if opponentMove == "paper" {
				pointsScored = 6 + 3
			} else {
				pointsScored = 6 + 1
			}
		} else if outcome == "lose" {
			if opponentMove == "rock" {
				pointsScored = 3
			} else if opponentMove == "paper" {
				pointsScored = 1
			} else {
				pointsScored = 2
			}
		}

		// part 1
		//opponentMove, myMove := letterConversion[splitLine[0]], letterConversion[splitLine[1]]
		//fmt.Print(opponentMove, " ", myMove, "\n")
		//if opponentMove == myMove {
		//	pointsScored = 3 + movePoints[myMove]
		//} else if myMove == "rock" && opponentMove == "scissors" ||
		//	myMove == "paper" && opponentMove == "rock" ||
		//	myMove == "scissors" && opponentMove == "paper" {
		//	pointsScored = 6 + movePoints[myMove]
		//} else if myMove == "rock" && opponentMove == "paper" ||
		//	myMove == "paper" && opponentMove == "scissors" ||
		//	myMove == "scissors" && opponentMove == "rock" {
		//	pointsScored = movePoints[myMove]
		//}


		totalPointsScored = totalPointsScored + pointsScored
		fmt.Print("PointsScored: ", pointsScored, "\n")
	}
	fmt.Print("totalPointsScored: ", totalPointsScored)
}
