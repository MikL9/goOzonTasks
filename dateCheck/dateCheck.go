package main

import "fmt"

func isLongYear(year int) bool {
	return (year%4 == 0 && year&100 != 0) || (year%400 == 0)
}

func checkDateValid(day int, month int, year int) bool {
	if year < 1950 || year > 2300 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}

	daysInMonth := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLongYear(year) {
		daysInMonth[2] = 29
	}

	return day <= daysInMonth[month]
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var day, month, year int
		fmt.Scan(&day, &month, &year)

		if checkDateValid(day, month, year) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
