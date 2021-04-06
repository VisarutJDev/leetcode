package datastructure

// Queue represent a queue that hold slice
type Queue struct {
	items []int
}

// Enqueue add value
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue remove value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
