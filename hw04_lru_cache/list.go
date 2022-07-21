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
	Value interface{} //значение
	next  *ListItem   //следующий элемент
	prev  *ListItem   //предыдущий элемент
	Elem  *list
}

type list struct {
	root ListItem
	len  int
}

// Nextious returns the next list element or nil.
func (e *ListItem) Nextious() *ListItem {
	if p := e.next; e.Elem != nil && p != &e.Elem.root {
		return p
	}
	return nil
}

// Previous returns the previous list element or nil.
func (e *ListItem) Previous() *ListItem {
	if p := e.prev; e.Elem != nil && p != &e.Elem.root {
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
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert inserts e after at, increments l.len, and returns e.
func (l *list) insert(e, at *ListItem) *ListItem {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
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
	if e == nil {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.Elem = nil
	l.len--
}

// Init initializes or clears list l.
func (l *list) Init() *list {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// lazyInit lazily initializes a zero List value.
func (l *list) lazyInit() {
	if l.root.next == nil {
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
	return l.insertValue(v, l.root.prev)
}

// move moves e to next to at.
func (l *list) Move(e, at *ListItem) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *list) MoveToFront(e *ListItem) {
	if e.Elem != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.Move(e, &l.root)
}

func NewList() *list {
	//return &list{}
	return new(list).Init()
}
