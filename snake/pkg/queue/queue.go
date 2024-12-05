package queue

type Queue[T any] struct {
	data        []T
	front, tail int
}

func New[T any](cap int) *Queue[T] {
	return &Queue[T]{
		data:  make([]T, 0, cap),
		front: 0,
		tail:  0,
	}
}

func (q *Queue[T]) Push(v T) {

}

func (q *Queue[T]) Pop() (T, bool) {
	return q.data[0], false
}

func (q *Queue[T]) Reset() {
	q.data = q.data[:]
	q.front = 0
	q.tail = 0
}

func (q *Queue[T]) ForEach(fn func(v T)) {

}
