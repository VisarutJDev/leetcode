package datastructure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	fmt.Println(stack)
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
