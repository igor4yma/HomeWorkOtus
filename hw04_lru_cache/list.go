package hw04lrucache

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
	front, back *ListItem
	len         int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = item
	} else {
		l.back = item
	}
	l.front = item
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v, Prev: l.back}
	if l.back != nil {
		l.back.Next = item
	} else {
		l.back = item
	}
	l.back = item
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil:
		l.front = i.Next
		i.Next.Prev = i.Prev
	case i.Next == nil:
		l.back = i.Prev
		i.Prev.Next = i.Next
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	l.Remove(i)
	i.Next = l.front
	l.front = i
}
