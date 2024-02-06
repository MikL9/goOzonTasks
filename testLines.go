package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	expectedOutputs := []string{"4", "TD", "JS"}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i >= len(expectedOutputs) {
			fmt.Printf("Extra line in file: %s\n", line)
			continue
		}
		if strings.TrimSpace(line) != expectedOutputs[i] {
			fmt.Printf("Mismatch! Expected: %s, Got: %s\n", expectedOutputs[i], line)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
