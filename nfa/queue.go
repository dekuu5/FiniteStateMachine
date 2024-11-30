package nfa

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{items: []interface{}{}}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Copy() *Queue {
	newQueue := NewQueue()
	newQueue.items = append(newQueue.items, q.items...)
	return newQueue
}

func (q *Queue) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}
