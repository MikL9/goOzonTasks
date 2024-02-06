package main

import (
	"bufio"
	"fmt"
	"os"
)

func processInput(input string) string {
	var currentLines [][]rune
	var currentLine *[]rune
	cursorPosition := 0
	linePosition := 0

	for _, char := range input {
		var stringLines [][]string
		for _, line := range currentLines {
			stringLine := make([]string, len(line))
			for i, char := range line {
				stringLine[i] = string(char)
			}
			stringLines = append(stringLines, stringLine)
		}
		var newLine []rune
		if linePosition >= len(currentLines) {
			currentLines = append(currentLines, newLine)
		}
		currentLine = &currentLines[linePosition]
		switch char {
		case 'L':
			if cursorPosition > 0 {
				cursorPosition--
			}
		case 'R':
			if cursorPosition < len(*currentLine) {
				cursorPosition++
			}
		case 'U':
			if linePosition != 0 {
				linePosition--
				currentLine = &currentLines[linePosition]
				if len(*currentLine) < cursorPosition {
					cursorPosition = len(*currentLine)
				}
			}
		case 'D':
			if linePosition+1 < len(currentLines) {
				linePosition++
				currentLine = &currentLines[linePosition]
				if len(*currentLine) < cursorPosition {
					cursorPosition = len(*currentLine)
				}
			}
		case 'B':
			cursorPosition = 0
		case 'E':
			if len(*currentLine) > 0 {
				cursorPosition = len(*currentLine)
			}
		case 'N':
			newLine = append([]rune{}, (*currentLine)[cursorPosition:]...)
			*currentLine = (*currentLine)[:cursorPosition]
			if linePosition+1 < len(currentLines) {
				upperSlice := append([][]rune{}, currentLines[:linePosition]...)
				upperSlice = append(upperSlice, *currentLine)
				upperSlice = append(upperSlice, []rune{})
				downSlice := currentLines[linePosition+1:]
				currentLines = append(upperSlice, downSlice...)
			} else {
				currentLines = append(currentLines, []rune{})
			}
			currentLines[linePosition+1] = newLine
			linePosition++
			cursorPosition = 0
		default:
			*currentLine = append((*currentLine)[:cursorPosition], append([]rune{char}, (*currentLine)[cursorPosition:]...)...)
			cursorPosition++
		}
	}

	var result []rune
	for _, line := range currentLines {
		fmt.Println(string(line))
	}

	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	numCases := 0
	fmt.Sscanf(t, "%d", &numCases)

	for i := 0; i < numCases; i++ {
		scanner.Scan()
		input := scanner.Text()

		processInput(input)

		fmt.Println("-")
	}
}
