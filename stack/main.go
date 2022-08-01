package main

import "fmt"

func main() {
	fmt.Println("hello stack!")
	fmt.Println(isValid("(()(())){}"))
}

type stack struct {
	data []rune
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
}

func (s *stack) pop() rune {
	l := len(s.data)
	if l == 0 {
		return rune(' ')
	}
	tmp := s.data[l-1]
	s.data = s.data[:l-1]
	return tmp
}

func isValid(s string) bool {
	st := stack{}
	for _, ch := range s {
		switch ch {
		case rune(']'):
			if st.pop() != rune('[') {
				return false
			}
		case rune(')'):
			if st.pop() != rune('(') {
				return false
			}
		case rune('}'):
			if st.pop() != rune('{') {
				return false
			}
		default:
			st.push(ch)
		}
	}
	return len(st.data) == 0
}
