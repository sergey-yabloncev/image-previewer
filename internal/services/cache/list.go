package cache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head, tail *ListItem
	count      int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.count++
	item := &ListItem{Value: v, Next: l.head}

	if l.head != nil {
		l.head.Prev = item
	} else {
		l.tail = item
	}

	l.head = item

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.count++

	item := &ListItem{Value: v, Prev: l.tail}

	if l.tail != nil {
		l.tail.Next = item
	} else {
		l.head = item
	}

	l.tail = item

	return item
}

func (l *list) Remove(i *ListItem) {
	l.count--
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	l.Remove(i)
	l.head.Prev = i
	i.Next = l.head
	l.head = i
}
