package main

import "fmt"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	if !s.checkItem(item) {
		s.items = append(s.items, item)
	}
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) checkItem(item interface{}) bool {
	for _, value := range s.items {
		if value == item {
			return true
		}
	}
	return false
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) values() []interface{} {
	return s.items
}

func main() {
	stack := Stack{}

	stack.Push("pisang")
	stack.Push("mangga")
	stack.Push("pisang")
	stack.Push("apel")
	stack.Push("nanas")
	stack.Push("jeruk")
	stack.Push("nanas")

	fmt.Println("Stack: ")
	fmt.Println(stack)
	fmt.Println("")

	fmt.Println("Pop: ")
	stack.Pop()
	fmt.Print("Stack after pop: ")
	fmt.Println(stack)

	fmt.Println("")
	fmt.Print("Stack is empty: ")
	fmt.Println(stack.isEmpty())

	fmt.Println("")
	fmt.Println("Stack values: ")
	fmt.Println(stack.values())
}
