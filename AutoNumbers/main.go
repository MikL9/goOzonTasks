package main

import (
	"fmt"
)

type PlateResult struct {
	IsValid bool
	Result  string
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var stringNumber string
		fmt.Scan(&stringNumber)

		resultNumber := validPlate(stringNumber)

		if resultNumber.IsValid == false {
			fmt.Println("-")
		} else {
			fmt.Println(resultNumber.Result)
		}
	}
}

func validPlate(plate string) PlateResult {
	var result PlateResult
	result.IsValid = true

	for {
		if len(plate) < 4 && len(plate) != 0 {
			result.IsValid = false
			break
		}
		if len(plate) == 0 {
			break
		}

		if isType1Plate(plate) {
			if result.Result != "" {
				result.Result += " "
			}
			result.Result += plate[:5]
			plate = plate[5:]
		} else if isType2Plate(plate) {
			if result.Result != "" {
				result.Result += " "
			}
			result.Result += plate[:4]
			plate = plate[4:]
		} else {
			result.IsValid = false
			break
		}
	}
	return result
}

func isType1Plate(plate string) bool {
	if len(plate) < 5 {
		return false
	}
	p := plate[:5]
	return isLetter(p[0]) && isNumeric(p[1]) && isNumeric(p[2]) &&
		isLetter(p[3]) && isLetter(p[4])
}

func isType2Plate(plate string) bool {
	p := plate[:4]
	return isLetter(p[0]) && isNumeric(p[1]) && isLetter(p[2]) &&
		isLetter(p[3])
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}
