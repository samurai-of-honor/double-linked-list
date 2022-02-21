package main

func main() {}

type Node struct {
	next, prev *Node
	list       *List
	data       string
}

type List struct {
	root Node
	len  int
}

// Init or clears list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns init list.
func New() *List {
	return new(List).Init()
}
