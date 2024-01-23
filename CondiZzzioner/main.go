package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		minTemp := 15
		maxTemp := 30

		for j := 0; j < n; j++ {
			scanner.Scan()
			condition := scanner.Text()

			temp, _ := strconv.Atoi(condition[3:])

			switch condition[0] {
			case '>':
				if temp > minTemp {
					minTemp = temp
				}
			case '<':
				if temp < maxTemp {
					maxTemp = temp
				}
			}
			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
