package datastructure

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}
	fmt.Println(queue)
	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
}
