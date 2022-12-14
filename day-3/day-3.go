package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getDiffCompartments(contents string) (string, string) {
	compartmentOne, compartmentTwo := contents[:len(contents)/2], contents[len(contents)/2:]
	return compartmentOne, compartmentTwo
}

func findCommonLetter(elfGroup [3]string) string {
	elfGroupArray := []string{elfGroup[0], elfGroup[1], elfGroup[2]}
	sort.Strings(elfGroupArray)
	sort.Slice(elfGroupArray, func(i, j int) bool {
		return len(elfGroupArray[i]) < len(elfGroupArray[j])
	})

	commonLetter := ""
	for letterIdx := 0; letterIdx < len(elfGroupArray[0]); letterIdx++ {
		if strings.Contains(elfGroupArray[2], string(elfGroupArray[0][letterIdx])) {
			// letter in last string
			if strings.Contains(elfGroupArray[1], string(elfGroupArray[0][letterIdx])) {
				// letter also in second string
				if !strings.Contains(commonLetter, string(elfGroupArray[0][letterIdx])) {
					commonLetter = commonLetter + string(elfGroupArray[0][letterIdx])
				}
			}
		}
	}
	return commonLetter
}

func getLetterScore(letter string) int {
	alphabetMapper := map[string]int{
		"a": 1,
	}
	count := 1
	for r := 'a'; r <= 'z'; r++ {
		alphabetMapper[string(r)] = count
		count = count + 1
	}

	letterScore := 0
	if letter == strings.ToUpper(letter) {
		letterScore = alphabetMapper[strings.ToLower(letter)] + 26
	} else {
		letterScore = alphabetMapper[letter]
	}
	return letterScore
}

func main() {
	file, _ := os.Open("day-3/rucksack-contents.txt")
	defer file.Close()

	totalPriorityCount := 0
	totalGroupCount := 0
	elfCount := 0
	var elfGroup [3]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		elfCount = elfCount + 1
		priorityCount := 0
		line := scanner.Text()
		elfGroup[elfCount-1] = line
		compartmentOne, compartmentTwo := getDiffCompartments(line)

		// looping through compartmentOne to check if those letters are in compartmentTwo
		nonUniqueLetters := ""
		for letterIdx := 0; letterIdx < len(compartmentOne); letterIdx++ {
			letter := string(compartmentOne[letterIdx])
			if strings.Contains(compartmentTwo, letter) {
				if !strings.Contains(nonUniqueLetters, letter) {
					nonUniqueLetters = nonUniqueLetters + letter
					priorityCount = priorityCount + getLetterScore(letter)
				}
			}
		}

		// splitting lines into group of three
		if elfCount == 3 {
			elfCount = 0

			// need to loop through the three contents and find letter in common
			commonLetter := findCommonLetter(elfGroup)
			commonLetterScore := getLetterScore(commonLetter)
			totalGroupCount = totalGroupCount + commonLetterScore

			var elfGroup []string = make([]string, 3)
			fmt.Print("elfGroup: ", elfGroup, "\n")
		}

		// adding to total count
		totalPriorityCount = totalPriorityCount + priorityCount
	}

	fmt.Print("total priority count: ", totalPriorityCount, "\n")
	fmt.Print("total group count: ", totalGroupCount, "\n")

}
