package main

import (
	s "patterns/star_patterns"
)

func main() {
	for i := 1; i < 40; i++ {
		s.PrintPattern(i)
	}
}

type Student struct {
	Name    string
	Age     int
	Address string
	School  School
}

type School struct {
	Name     string
	Students []Student
}
