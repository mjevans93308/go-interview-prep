package datastructures

type Stack struct {
	list List
}

func (s *Stack) SizeOf() int {
	return s.list.SizeOf()
}

func (s *Stack) Push(data int) {
	s.list.InsertAtFront(data)
}

func (s *Stack) Pop() (int, error) {
	data, err := s.list.DeleteFromFront()
	if err != nil {
		return -1, err
	}
	return data, nil
}

func (s *Stack) Peek() (int, error) {
	data, err := s.list.CheckAtIndex(0)
	if err != nil {
		return -1, err
	}
	return data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}
