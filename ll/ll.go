package ll

import (
	"bytes"
	"strconv"
)

// ConsCell is the struct that holds each cell of our linked list.
type ConsCell struct {
	Car int
	Cdr *ConsCell
}

// NIL is our sentinel for the end of the linked list. It also represents a
// linked list of 0 length.
var NIL = (*ConsCell)(nil)

// Cons creates a new cell to hold data, puts it at the head of the list already
// represented by c, and then returns that cell.
func (c *ConsCell) Cons(data int) *ConsCell {
	return &ConsCell{Car: data, Cdr: c}
}

// Len returns the length of the linked list.
func (c *ConsCell) Len() int {
	l := 0
	for current := c; current != NIL; current = current.Cdr {
		l++
	}

	return l
}

// String returns the string representation of this linked list.
func (c *ConsCell) String() string {
	out := &bytes.Buffer{}
	out.WriteString("(")
	for current := c; current != NIL; current = current.Cdr {
		if c != current {
			out.WriteString(", ")
		}
		out.WriteString(strconv.Itoa(current.Car))

	}
	out.WriteString(")")

	return out.String()
}

// Reverse reverses the linked list pointed to by c in place, and returns a
// pointer to the new head.
func (c *ConsCell) Reverse() *ConsCell {
	// oldPrev is what was previous to the current node.
	oldPrev := NIL
	// oldNext is the head of what we need to process next.
	oldNext := NIL

	for current := c; current != NIL; current = oldNext {
		oldNext = current.Cdr
		current.Cdr = oldPrev
		oldPrev = current
	}
	return oldPrev
}

// ReverseNextN reverses the first n elements of the list pointed to by c. The
// remaining list is detached. The return values of this function are 1) the new
// first element, 2) the new nth element, and 3) the detached list.
func (c *ConsCell) ReverseFirstN(n int) (first, nth, rest *ConsCell) {
	oldPrev := NIL
	oldNext := NIL

	current := c

	for ; current != NIL && n != 0; current, n = oldNext, n-1 {
		oldNext = current.Cdr
		current.Cdr = oldPrev
		oldPrev = current
	}

	return oldPrev, c, current
}

// ReverseBy3s returns the list pointed to by c changed in place so that every
// group of 3 is reversed.
func (c *ConsCell) ReverseBy3s() *ConsCell {
	newHead, tail, rest := c.ReverseFirstN(3)

	for rest != NIL {
		oldTail := tail
		oldTail.Cdr, tail, rest = rest.ReverseFirstN(3)
	}
	return newHead
}
