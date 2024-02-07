package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	dir = dir + "/Tree/testTree/50"
	fmt.Println("Current directory is:", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var expectedOutputs []string
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			contentStr := string(content)

			if strings.Contains(file.Name(), "a") {
				scanner := bufio.NewScanner(strings.NewReader(contentStr))
				i := 0
				for scanner.Scan() {
					line := scanner.Text()
					if i >= len(expectedOutputs) {
						fmt.Printf("Extra line in file%s: %s\n", file.Name(), line)
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

			} else {
				//expectedOutputs = mainFunc(contentStr)
			}
		}
	}
}
