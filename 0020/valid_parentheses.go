func isValid(s string) bool {
	a := []byte(s)
	stack := new(Stack)
	for i := 0; i < len(a); i++ {
		if v, err := stack.Top(); err != nil {
			stack.Push(a[i])
		} else {
			if a[i]-v < 3 && a[i] != v {
				stack.Pop()
			} else {
				stack.Push(a[i])
			}
		}
	}

	if _, err := stack.Top(); err != nil {
		return true
	}

	return false
}

type Stack []byte

func (stack *Stack) Push(value byte) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (byte, error) {
	if len(stack) == 0 {
		return 0, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (byte, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return 0, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
