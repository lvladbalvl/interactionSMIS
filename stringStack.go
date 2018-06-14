package main
type Stack struct {
	data []interface{}
}

// Push adds x to the top of the stack.
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It’s a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

// Prefix I indicates interface, not struct.
type IStack interface {
	Push(interface{})
	Pop() interface{}
	Size() int
}

// Here we specialize the Stack to require string elements.
// Compare with Java: Stack<String> s = new Stack<String>();

type StringStack struct {
	Stack
}

func (s *StringStack) Push(n string) { s.Stack.Push(n) }

// Laughs 'n' giggles. struct, not interface embedding, so infinite recursion here.
//func (s *StringStack) Push(n string) { s.Push(n) }

func (s *StringStack) Pop() string { return s.Stack.Pop().(string) }

// Unnecessary wrapper since string type parameter never used.
// func (s *StringStack) Size() int   { return s.Stack.Size() }

type StringIStack struct {
	IStack
}

// Necessary to force string input:
func (s *StringIStack) Push(n string) { s.IStack.Push(n) }

// Necessary to force string output:
func (s *StringIStack) Pop() string { return s.IStack.Pop().(string) }

// Unnecessary if value assigned to s already has Size() int method.
func (s *StringIStack) Size() int { return s.IStack.Size() }

type IStringStack interface {
	Push(string)
	Pop() string
	Size() int
}
