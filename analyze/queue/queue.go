package queue

// interface{} the type of the queue
type Queue struct {
	items []interface{}
}

// new Queue storing items
func NewQueue(size int) (q *Queue) {
	q = new(Queue)
	q.items = make([]interface{}, size)
	return q
}

// push item
func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

// pop item
func (q *Queue) Pop() (item interface{}) {
	item = q.items[0]
	q.items = q.items[1:]
	return
}

// the size of queue
func (q *Queue) Size() int {
	return len(q.items)
}

// whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
