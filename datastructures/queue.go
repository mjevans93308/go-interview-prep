package datastructures

type Queue struct {
	list List
}

func (q *Queue) SizeOf() int {
	return q.list.SizeOf()
}

func (q *Queue) Enqueue(data int) {
	q.list.InsertAtBack(data)
}

func (q *Queue) Dequeue() (int, error) {
	data, err := q.list.DeleteFromBack()
	if err != nil {
		return -1, err
	}
	return data, nil
}

func (q *Queue) Peek() (int, error) {
	oldest := q.SizeOf() - 1
	data, err := q.list.CheckAtIndex(oldest)
	if err != nil {
		return -1, err
	}
	return data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}
