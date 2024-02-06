package main

import (
	"fmt"
)

func compressSequence(n int, a []int) (int, []int) {
	result := make([]int, 0)
	i := 0

	for i < n {
		start := i
		action := true
		if i < n-1 && (a[i]+1 == a[i+1]) { // true +; false -;
			action = true
		} else if i < n-1 && (a[i]-1 == a[i+1]) {
			action = false
		}
		for i < n-1 && (((a[i]+1 == a[i+1]) && action == true) || ((a[i]-1 == a[i+1]) && action == false)) {
			i++
		}
		result = append(result, a[start], a[i]-a[start])
		i++
	}

	return len(result), result
}

func main() {
	var t, n int
	fmt.Scan(&t)

	for test := 0; test < t; test++ {
		fmt.Scan(&n)
		a := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		length, compressed := compressSequence(n, a)

		fmt.Println(length)
		for i := 0; i < len(compressed); i++ {
			fmt.Print(compressed[i], " ")
		}
		fmt.Println()
	}
}
