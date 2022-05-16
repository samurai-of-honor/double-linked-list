package main

type Node struct {
	next, prev *Node
	list       *List
	data       string
}

type List struct {
	root Node
	len  int
}

// NewList create new init list
func NewList() *List {
	return new(List)
}

// Length returns list length
func (l *List) Length() int {
	return l.len
}

// Clear cleaning list
func (l *List) Clear() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// First returns the first node of list
func (l *List) First() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Next returns the next list node
func (n *Node) Next() *Node {
	fol := n.next
	if n.list != nil && fol != &n.list.root {
		return fol
	} else {
		return nil
	}
}

// Display returns slice with all values of list
func (l *List) Display() []string {
	sl := make([]string, 0)
	for n := l.First(); n != nil; n = n.Next() {
		sl = append(sl, n.data)
	}
	return sl
}

// Get returns node by its index
func (l *List) Get(i int) *Node {
	if i < 0 {
		panic("index < 0")
	}
	n := l.First()
	for ; i > 0; i-- {
		n = n.Next()
	}
	return n
}

// Add inserts value v after at
func (l *List) add(v string, at *Node) *Node {
	n := &Node{data: v}
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
	return n
}

// Insert inserts a new node with value v at given index
func (l *List) Insert(v string, i int) *Node {
	at := l.Get(i)
	if at == nil {
		l.Clear()
		return l.add(v, l.root.prev)
	} else {
		return l.add(v, at.prev)
	}
}

// Append inserts a new node with value v at the back of list
func (l *List) Append(v string) *Node {
	if l.root.next == nil {
		l.Clear()
	}
	return l.add(v, l.root.prev)
}

// Delete removes node from list by index
func (l *List) Delete(i int) string {
	n := l.Get(i)
	if n.list == l {
		n.prev.next = n.next
		n.next.prev = n.prev
		n.next = nil
		n.prev = nil
		n.list = nil
		l.len--
	}
	return n.data
}
