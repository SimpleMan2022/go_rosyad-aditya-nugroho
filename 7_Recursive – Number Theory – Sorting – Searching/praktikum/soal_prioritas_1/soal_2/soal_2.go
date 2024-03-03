package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
	// output:
	// best student in math: aziz with 100
	// best student in science: jamie with 88
	// best student in english: jamie with 90
	// best student in average: aziz with 84

}

func bestMath(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].MathScore > students[j].MathScore
	})
}

func bestScience(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].ScienceScore > students[j].ScienceScore
	})
}

func bestEnglish(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].EnglishScore > students[j].EnglishScore
	})
}

func bestAverage(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].MathScore+students[i].ScienceScore+students[i].EnglishScore >
			students[j].MathScore+students[j].ScienceScore+students[j].EnglishScore
	})
}

func getInfo(students []Student) {
	bestMath(students)
	fmt.Println("best student in math:", students[0].Name, "with", students[0].MathScore)
	bestScience(students)
	fmt.Println("best student in science:", students[0].Name, "with", students[0].ScienceScore)
	bestEnglish(students)
	fmt.Println("best student in english:", students[0].Name, "with", students[0].EnglishScore)
	bestAverage(students)
	fmt.Println("best student in average:", students[0].Name, "with", (students[0].MathScore+students[0].ScienceScore+students[0].EnglishScore)/3)
}
