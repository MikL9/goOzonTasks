package main

import (
	"fmt"
	"strconv"
	"strings"
)

func buildRes(res []string, r1 int, r2 int) []string {
	if r1 >= r2 {
		return append(res, strconv.Itoa(r1))
	} else {
		return append(res, fmt.Sprintf("%d-%d", r1, r2))
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)
		var line string
		fmt.Scan(&line)

		printed := make([]bool, k+1)
		printedLine := strings.Split(line, ",")

		for _, page := range printedLine {
			if strings.Contains(page, "-") {
				rangedPages := strings.Split(page, "-")
				first, _ := strconv.Atoi(rangedPages[0])
				last, _ := strconv.Atoi(rangedPages[1])

				for j := first; j <= last; j++ {
					printed[j] = true
				}
			} else {
				{
					page, _ := strconv.Atoi(page)
					printed[page] = true
				}
			}
		}
		var result []string
		var range1 int
		var range2 int
		for j := 1; j <= k; j++ {
			if printed[j] == false {
				if j+1 <= k {
					if printed[j+1] == false {
						range2 = j + 1
					}
				}
				if range1 >= range2 || range1 == 0 {
					range1 = j
				}
				if j == k {
					result = buildRes(result, range1, range2)
				}
				if j+1 <= k {
					if printed[j+1] == true {
						result = buildRes(result, range1, range2)
					}
				}
			} else if j != 1 && range1 != 0 {
				if j <= k {
					range1 = j + 1
					range2 = range1
				}
			}
		}

		fmt.Println(strings.Join(result, ","))
	}
}
