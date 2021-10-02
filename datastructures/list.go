package datastructures

import "errors"

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func (l *List) IsEmpty() bool {
	return l.length == 0
}

func (l *List) InsertAtFront(data int) {
	n := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if l.IsEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.head.prev = n
		n.next = l.head
		l.head = n
	}
}

func (l *List) InsertAtBack(data int) {
	n := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if l.IsEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
}

func (l *List) DeleteFromFront() (int, error) {
	if l.IsEmpty() {
		return -1, errors.New("list empty, nothing to delete")
	} else {
		data := l.head.data

		// list has 1 item
		if l.length == 1 {
			return data, nil
		}

		curr := l.head
		l.head = curr.next
		l.head.prev = nil

		l.length -= 1
		return data, nil
	}
}

func (l *List) DeleteFromBack() (int, error) {
	if l.IsEmpty() {
		return -1, errors.New("list empty, nothing to delete")
	} else {
		data := l.tail.data

		// list has 1 item
		if l.length == 1 {
			return data, nil
		}

		curr := l.tail
		l.tail = curr.prev
		l.tail.next = nil

		return data, nil
	}
}

func (l *List) CheckAtIndex(pos int) (int, error) {
	if l.IsEmpty() {
		return -1, errors.New("cannot fetch item from empty list")
	}

	if l.SizeOf() <= pos {
		return -1, errors.New("not enough items in list to satisfy pos arg")
	}

	curr := l.head
	index := 0
	for curr.next != nil && index <= pos {
		curr = curr.next
		index += 1
	}
	return curr.data, nil
}

func (l *List) SizeOf() int {
	// even if empty, length should be initialized to 0
	return l.length
}
