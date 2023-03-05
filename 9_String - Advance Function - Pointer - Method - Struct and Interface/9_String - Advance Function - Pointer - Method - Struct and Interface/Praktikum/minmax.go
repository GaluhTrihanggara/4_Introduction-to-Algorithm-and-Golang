// Program Untuk Mencari Nilai Minimal Dan Maksimal
// by Galuh Trihanggara
package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func (s *Student) GetAverageScore() float64 {
	sum := 0
	for _, score := range s.Score {
		sum += score
	}

	return float64(sum) / float64(len(s.Score))
}

func (s *Student) GetMinimumScore() int {
	min := math.MaxInt32
	for _, score := range s.Score {
		if score < min {
			min = score
		}
	}

	return min
}

func (s *Student) GetMaximumScore() int {
	max := math.MinInt32
	for _, score := range s.Score {
		if score > max {
			max = score
		}
	}

	return max
}

func main() {
	students := []Student{
		{Name: "Rizki", Score: []int{80}},
		{Name: "Kobar", Score: []int{75}},
		{Name: "Ismail", Score: []int{70}},
		{Name: "Umam", Score: []int{75}},
		{Name: "Yopan", Score: []int{60}},
	}

	var sum float64
	var min, max int
	for i, student := range students {
		sum += student.GetAverageScore()

		if i == 0 {
			min = student.GetMinimumScore()
			max = student.GetMaximumScore()
		} else {
			if student.GetMinimumScore() < min {
				min = student.GetMinimumScore()
			}
			if student.GetMaximumScore() > max {
				max = student.GetMaximumScore()
			}
		}
	}

	avg := sum / float64(len(students))

	fmt.Printf("Average Score : %.0f\n", avg)

	for _, student := range students {
		if student.GetMinimumScore() == min {
			fmt.Printf("Min Score of Students : %s (%d)\n", student.Name, min)
		}

		if student.GetMaximumScore() == max {
			fmt.Printf("Max Score of Students : %s (%d)\n", student.Name, max)
		}
	}
}
