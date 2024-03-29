package main

import (
	"fmt"
)

func uniquePathsWithObstacles(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	if grid[x-1][y-1] == 1 || grid[0][0] == 1 {
		return 0
	}
	for i := x - 1; i >= 0; i-- {
		for j := y - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				continue
			}
			if i == x-1 && j == y-1 {
				grid[i][j] = 1
			} else {
				if j < y-1 {
					grid[i][j] += grid[i][j+1]
				}
				if i < x-1 {
					grid[i][j] += grid[i+1][j]
				}
			}
		}
	}

	return grid[0][0]
}

func main() {
	grid := [][]int{
		{
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			1,
			1,
			0,
			0,
			1,
			0,
			1,
			1,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
		},
		{
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
		},
		{
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
		},
		{
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
		},
		{
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			1,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
		},
		{
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
			1,
			0,
			0,
			0,
			1,
		},
		{
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			1,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
		},
		{
			1,
			0,
			1,
			0,
			1,
			1,
			0,
			1,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
		},
		{
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
		},
		{
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
		},
		{
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			1,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
		},
		{
			0,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			1,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			1,
			0,
			1,
			1,
			1,
			0,
			0,
			1,
			0,
			1,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		{
			1,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
		},
		{
			0,
			1,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			1,
			0,
			1,
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			1,
			1,
			0,
			0,
		},
		{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			1,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
