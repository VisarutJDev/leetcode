package easy

import (
	"fmt"
	"testing"
)

func TestCheckTwoChessboards(t *testing.T) {
	fmt.Println(checkTwoChessboards("a1", "c3"))
	fmt.Println(checkTwoChessboards("a1", "h3"))
}
