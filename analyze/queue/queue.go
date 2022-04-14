package queue

// interface{} the type of the queue
type Queue struct {
	items []interface{}
}

// new Queue storing items
func NewQueue() (q *Queue) {
	q = new(Queue)
	q.items = make([]interface{}, 0)
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

// the front item
func (q *Queue) Front() (item interface{}) {
	return q.items[0]
}

// the size of queue
func (q *Queue) Size() int {
	return len(q.items)
}

// whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
