package MoskoiBoy

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsValidShipCounts(t *testing.T) {
	testCases := []struct {
		shipSizes []int
		expected  bool
	}{
		{[]int{2, 1, 3, 1, 2, 3, 1, 1, 4, 2}, true},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}, false},
		{[]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4}, true},
		{[]int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}, true},
		{[]int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, false},
	}

	for _, tc := range testCases {
		t.Run(strings.Join(strings.Fields(fmt.Sprint(tc.shipSizes)), "_"), func(t *testing.T) {
			result := isValidShipCounts(tc.shipSizes)
			if result != tc.expected {
				t.Errorf("For shipSizes %v, got %v, expected %v", tc.shipSizes, result, tc.expected)
			}
		})
	}
}
