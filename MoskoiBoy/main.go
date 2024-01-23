package MoskoiBoy

import "fmt"

func isValidShipCounts(shipSizes []int) bool {
	counts := map[int]int{
		1: 4, // однопалубные корабли
		2: 3, // двухпалубные корабли
		3: 2, // трехпалубные корабли
		4: 1, // четырехпалубные корабли
	}

	for _, size := range shipSizes {
		if counts[size] == 0 {
			return false
		}
		counts[size]--
	}

	return true
}

func ds() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		shipSizes := make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Scan(&shipSizes[j])
		}

		if isValidShipCounts(shipSizes) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
