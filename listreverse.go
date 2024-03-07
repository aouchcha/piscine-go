package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	l.Tail = l.Head
	var avant *NodeL
	for l.Head != nil {
		next := l.Head.Next
		l.Head.Next = avant
		avant = l.Head
		l.Head = next
	}
	l.Head = avant
	l.Tail.Next = nil
}
