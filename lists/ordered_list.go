package lists

import (
	"n-puzzle-problem/nodes"
	"n-puzzle-problem/utils"
)

type item[T any] struct {
	val   any
	next  *item[T]
	index uint64
}

type List[T any] struct {
	head *item[T]
}

func (l *List[T]) Insert(val T, index uint64) {
	item := item[T]{val, nil, index}

	// fmt.Println("inserted val", val)
	if l.head == nil {
		l.head = &item
		// fmt.Println("l.head", l.head)
		return
	}
	// fmt.Println("AAAAAA")
	ptr := l.head
	if item.index < ptr.index {
		l.head = &item
		item.next = ptr
		return
	}

	for ptr.next != nil && ptr.next.index < item.index {
		ptr = ptr.next
	}

	item.next = ptr.next
	ptr.next = &item
}

func (l *List[T]) RemoveFirst() {
	if l.head != nil && l.head.next != nil {
		l.head = l.head.next
	} else {
		l.head = nil
	}
}

func (l *List[T]) GetFirst() T {
	// fmt.Println(l.head)

	if l.head.val != nil {
		if v, ok := l.head.val.(T); ok {
			return v
		}
	}

	return getZero[T]()
}

func (l *List[T]) ForEach(callback func(node nodes.Node)) {
	ptr := l.head
	for ptr != nil {
		if ptr.val != nil {
			if v, ok := ptr.val.(nodes.Node); ok {
				callback(v)
			}
		}
		ptr = ptr.next
	}
}

func getZero[T any]() T {
	var result T
	return result
}

func (l *List[T]) Contains(c utils.Comparator) bool {
	contains := false
	l.ForEach(func(node nodes.Node) {
		if c.Equals(&node) {
			contains = true
		}
	})

	return contains
}

func (l *List[T]) Length() int {
	length := 0
	l.ForEach(func(node nodes.Node) {
		length += 1
	})

	return length
}
