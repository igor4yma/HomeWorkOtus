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
	Elem  *list
}

type list struct {
	List
	root ListItem
	len  int
}

// Nextious returns the next list element or nil.
func (e *ListItem) Nextious() *ListItem {
	if p := e.Next; e.Elem != nil && p != &e.Elem.root {
		return p
	}
	return nil
}

// Previous returns the previous list element or nil.
func (e *ListItem) Previous() *ListItem {
	if p := e.Prev; e.Elem != nil && p != &e.Elem.root {
		return p
	}
	return nil
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *list) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Prev
}

// insert inserts e after at, increments l.len, and returns e.
func (l *list) insert(e, at *ListItem) *ListItem {
	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
	e.Elem = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *list) insertValue(v interface{}, at *ListItem) *ListItem {
	return l.insert(&ListItem{Value: v}, at)
}

// Remove removes e from its list, decrements l.len
func (l *list) Remove(e *ListItem) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil // avoid memory leaks
	e.Prev = nil // avoid memory leaks
	e.Elem = nil
	l.len--
}

// Init initializes or clears list l.
func (l *list) Init() *list {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.len = 0
	return l
}

// lazyInit lazily initializes a zero List value.
func (l *list) lazyInit() {
	if l.root.Next == nil {
		l.Init()
	}
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *list) PushFront(v interface{}) *ListItem {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *list) PushBack(v interface{}) *ListItem {
	l.lazyInit()
	return l.insertValue(v, l.root.Prev)
}

// move moves e to next to at.
func (l *list) Move(e, at *ListItem) {
	if e == at {
		return
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev

	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *list) MoveToFront(e *ListItem) {
	if e.Elem != l || l.root.Next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.Move(e, &l.root)
}

func NewList() List {
	return new(list)
}
