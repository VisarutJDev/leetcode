package hackerrank

import (
	"fmt"
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {
	// Write your code here
	mm := DesignerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc")
	fmt.Println(mm)
}
