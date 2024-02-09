package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func dfs(x, y int, grid [][]byte, visited [][]bool, depth int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || visited[x][y] || grid[x][y] == '.' {
		return depth
	}
	visited[x][y] = true
	return max(
		dfs(x-1, y, grid, visited, depth+1),
		dfs(x+1, y, grid, visited, depth+1),
		dfs(x, y-1, grid, visited, depth+1),
		dfs(x, y+1, grid, visited, depth+1),
	)
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func solve() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	grid := make([][]byte, n)
	visited := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]byte, m)
		visited[i] = make([]bool, m)
		fmt.Fscan(in, &grid[i])
	}
	var res []int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '*' && !visited[i][j] {
				res = append(res, dfs(i, j, grid, visited, 0))
			}
		}
	}
	sort.Ints(res)
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		solve()
		t--
	}
}
