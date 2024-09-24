package hackerrank

import (
	"fmt"
	"testing"
)

func TestBomberMan(t *testing.T) {
	// fmt.Println("Result >> ", bomberMan(3, []string{
	// 	".......", "...O...", "....O..", ".......", "OO.....", "OO.....",
	// }))

	fmt.Println("Result >> ", bomberMan(5, []string{
		".......",
		"...O.O.",
		"....O..",
		"..O....",
		"OO...OO",
		"OO.O...",
	}))
}
