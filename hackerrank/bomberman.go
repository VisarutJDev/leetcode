package hackerrank

import (
	"fmt"
	"strings"
)

func bomberMan(n int32, grid []string) []string {
	// Write your code here
	// n is second
	// problem define that's
	// 0 sec >> bomb plant
	// 1 sec >> nothing countdown
	// 2 sec >> plant bomb to empty space
	// 3 sec >> bomb that already plant in initial phase explode >> bomb from "2 sec" do nothing countdown
	// 4 sec >> plant bomb to empty space
	// 5 sec >> bomb from "2 sec" explode >> bomb from "4 sec" do nothing countdown
	// 6 sec >> plant bomb to empty space
	// 7 sec >> bomb from "4 sec" explode

	// even number is bomb plant
	// odd number is bomb from "x (even) second" ago explode and bomb just plant second ago do nothing

	// get init bomb position
	for i := int32(0); i <= n; i++ {
		for j := 0; j < len(grid); j++ {
			for strings.Contains(grid[j], "1") {
				k := strings.Index(grid[j], "1")
				//boooom at j, strings.Index(grid[j], "1")
				grid[j] = strings.Replace(grid[j], "1", ".", 1)
				//boooom at j, strings.Index(grid[j], "1") + 1
				//boooom at j, strings.Index(grid[j], "1") - 1
				s := []byte(grid[j])
				if k != len(grid[j])-1 {
					if string(s[k+1]) != "1" {
						s[k+1] = '.'
					}
				}
				if k != 0 {
					if string(s[k-1]) != "1" {
						s[k-1] = '.'
					}
				}
				grid[j] = string(s)
				//boooom at j - 1, strings.Index(grid[j], "1")
				//boooom at j + 1, strings.Index(grid[j], "1")
				if j != len(grid)-1 {
					x := []byte(grid[j+1])
					if string(x[k]) != "1" {
						x[k] = '.'
						grid[j+1] = string(x)
					}

				}
				if j != 0 {
					x := []byte(grid[j-1])
					if string(x[k]) != "1" {
						x[k] = '.'
						grid[j-1] = string(x)
					}

				}
			}

			grid[j] = strings.ReplaceAll(grid[j], "2", "1")
			grid[j] = strings.ReplaceAll(grid[j], "3", "2")
			grid[j] = strings.ReplaceAll(grid[j], "O", "3")
			if i%2 == 0 {
				grid[j] = strings.ReplaceAll(grid[j], ".", "O")
			}

			// reset
			if i == n {
				grid[j] = strings.ReplaceAll(grid[j], "3", "O")
				grid[j] = strings.ReplaceAll(grid[j], "2", "O")
				grid[j] = strings.ReplaceAll(grid[j], "1", "O")
			}
		}
		fmt.Println(grid)
	}

	return grid
}
