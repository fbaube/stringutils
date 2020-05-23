package stringutils

// Stringstack is a LIFO stack for strings.
type Stringstack []string

// func NewStringStack() { return stringstack }

// IsEmpty is a no-brainer.
func (ss Stringstack) IsEmpty() bool { return len(ss) == 0 }

// Peek returns empty string ("") on an empty stack.
func (ss Stringstack) Peek() string {
	if ss.IsEmpty() { return "" }
	return ss[len(ss)-1]
}

// Push might reslice the stack.
func (ss *Stringstack) Push(s string) { (*ss) = append((*ss), s) }

// Pop returns empty string ("") on an empty stack.
func (ss *Stringstack) Pop() string {
	if ss.IsEmpty() { return "" }
	d := (*ss)[len(*ss)-1]
	(*ss) = (*ss)[:len(*ss)-1]
	return d
}
