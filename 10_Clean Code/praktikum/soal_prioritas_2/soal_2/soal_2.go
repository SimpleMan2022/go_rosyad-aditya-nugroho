package main

import "fmt"

type Set struct {
	data map[int]bool
}

func NewSet() *Set {
	return &Set{
		data: map[int]bool{},
	}
}

func (s *Set) Add(i int) {
	s.data[i] = true
}

func (s *Set) Get() []int {
	result := []int{}

	for key := range s.data {
		result = append(result, key)
	}

	return result
}

func (s *Set) Remove(key int) {
	delete(s.data, key)
}

func main() {
	s := NewSet()

	s.Add(1)
	s.Add(2)
	s.Add(1)
	s.Add(3)
	s.Add(4)
	s.Add(5)

	s.Remove(100)

	fmt.Println(s.Get())
}
