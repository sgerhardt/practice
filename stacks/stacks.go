package stacks

type StackWithMax struct {
	Stack []int
	Max   int
}

func (s *StackWithMax) push(i int) {
	if i > s.Max {
		s.Max = i
	}
	s.Stack = append(s.Stack, i)
}

// Remove and return top element of stack. Return empty 0 if stack is empty.
func (s *StackWithMax) pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(s.Stack) - 1   // Get the index of the top most element.
		element := (s.Stack)[index] // Index into the slice and obtain the element.
		s.Stack = (s.Stack)[:index] // Remove it from the stack by slicing it off.
		if element == s.Max {
			s.Max = s.peek()
		}
		return element
	}
}

func (s *StackWithMax) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *StackWithMax) peek() int {
	if s == nil || s.IsEmpty() {
		return 0
	}
	index := len(s.Stack) - 1 // Get the index of the top most element.
	return (s.Stack)[index]
}

func (s *StackWithMax) GetMax() int {
	//largestSoFar := 0
	//for _, val := range s.StackInt {
	//	if val > largestSoFar {
	//		largestSoFar = val
	//	}
	//}

	return s.Max
}

type StackInt struct {
	Stack []int
}

func (s *StackInt) push(i int) {
	s.Stack = append(s.Stack, i)
}

// Remove and return top element of stack. Return empty 0 if stack is empty.
func (s *StackInt) pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(s.Stack) - 1   // Get the index of the top most element.
		element := (s.Stack)[index] // Index into the slice and obtain the element.
		s.Stack = (s.Stack)[:index] // Remove it from the stack by slicing it off.
		return element
	}
}

func (s *StackInt) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *StackInt) peek() int {
	if s == nil || s.IsEmpty() {
		return 0
	}
	index := len(s.Stack) - 1 // Get the index of the top most element.
	return (s.Stack)[index]
}

type StackString struct {
	Stack []string
}

func (s *StackString) push(i string) {
	s.Stack = append(s.Stack, i)
}

// Remove and return top element of stack. Return empty "" if stack is empty.
func (s *StackString) pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(s.Stack) - 1   // Get the index of the top most element.
		element := (s.Stack)[index] // Index into the slice and obtain the element.
		s.Stack = (s.Stack)[:index] // Remove it from the stack by slicing it off.
		return element
	}
}

func (s *StackString) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *StackString) peek() string {
	if s == nil || s.IsEmpty() {
		return ""
	}
	index := len(s.Stack) - 1 // Get the index of the top most element.
	return (s.Stack)[index]
}
