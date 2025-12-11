package educativedsa

import "fmt"

type AnyStack struct {
	data []interface{}
}

// IsEmpty

func (s *AnyStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Length

func (s *AnyStack) Length() int {
	return len(s.data)
}

// Pop

func (s *AnyStack) Pop() (value interface{}) {

	if s.IsEmpty() {
		return nil
	}

	length := s.Length()

	value = s.data[length-1]

	s.data = s.data[:length-1]

	return value
}

// Peek

func (s *AnyStack) Peek() (value interface{}) {
	if !s.IsEmpty() {
		length := s.Length()

		value = s.data[length-1]

		return value
	}

	return nil
}

// Push

func (s *AnyStack) Push(value interface{}) {
	s.data = append(s.data, value)
}

func LearnedBalancedStrings(value string) bool {

	if len(value) <= 1 {
		return false
	}

	var (
		stack1 AnyStack
		i      = 0
	)

	for i < len(value) {

		stack1.Push(string(value[i]))

		if string(value[i]) == "}" {
			stackValue := stack1.Peek()
			if stackValue != nil && stackValue.(string) == "{" {
				stack1.Pop()
			}
		}

		if string(value[i]) == ")" {
			stackValue := stack1.Peek()
			if stackValue != nil && stackValue.(string) == "(" {
				stack1.Pop()
			}
		}

		if string(value[i]) == "]" {
			stackValue := stack1.Peek()
			if stackValue != nil && stackValue.(string) == "[" {
				stack1.Pop()
			}
		}

		i++
	}
	fmt.Println("ty", stack1.data)

	return stack1.IsEmpty()

}

func BalancedStrings(s string) bool {
	stk := new(AnyStack)

	for _, char := range s {
		switch char {
		case '{', '(', '[':
			stk.Push(char)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		}
	}

	return stk.IsEmpty()
}

func MaxDepthParanthesis(exp string) int {
	var (
		depth, maxDepth int
		stringStack     AnyStack
	)

	for _, char := range exp {

		if char == '(' {

			// fmt.Println("maxDepss", maxDepth, depth)

			if depth != 0 && maxDepth != 0 {

				depth--
				for stringStack.Length() != 0 {
					stringStack.Pop()
				}
			}
			stringStack.Push(char)
		}

		if char == ')' {
			stringStack.Pop()
			depth++

		}

		maxDepth = depth

	}

	return maxDepth
}

func LongestValidParanthesis(exp string) int {
	newStack := new(AnyStack)

	length := 0

	newStack.Push(-1)

	for index, char := range exp {
		if char == '(' {
			newStack.Push(index)
		} else {
			newStack.Pop()

			if newStack.Length() != 0 {
				if length < index-newStack.Peek().(int) {
					length = index - newStack.Peek().(int)
				}

			} else {
				newStack.Push(index)
			}
		}
	}

	return length
}
