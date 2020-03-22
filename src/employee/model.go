package employee

import (
	"fmt"
	"time"
)

type Employee struct {
	Name      string
	BirthDate time.Time
	Stats     Stats
	Salaries  []Salary
}

type Stats struct {
	Average float32
	Min     float32
	Max     float32
	Sum     float32
}

type Salary struct {
	Date   time.Time
	Salary float32
	Bonus  float32
}

func printSlice(s []Salary) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printEmployee(employee Employee) {
	fmt.Printf("Name: %v\n", employee.Name)
	fmt.Printf("Salaries\n")
	layout := "2006-01-02"
	for _, v := range employee.Salaries {
		fmt.Printf("\tDate: %v. Salary %v (%v as bonus)\n", v.Date.Format(layout), v.Salary, v.Bonus)
	}
	employee.initStats()
}

func (employee *Employee) initStats() {
	var min float32 = 0
	var max float32 = 0
	var sum float32 = 0
	var average float32 = 0
	for _, v := range employee.Salaries {
		sum = sum + v.Salary
		if min == 0 && max == 0 {
			min = v.Salary
			max = v.Salary
		}
		if v.Salary < min {
			min = v.Salary
		}
		if v.Salary > max {
			max = v.Salary
		}
	}
	average = sum / float32(len(employee.Salaries))
	employee.Stats.Average = average
	employee.Stats.Max = max
	employee.Stats.Min = min
	employee.Stats.Sum = sum
	//fmt.Printf("Stats summary\n")
	//fmt.Printf("\tSum: %v\n", sum)
	//fmt.Printf("\tMin: %v\n", min)
	//fmt.Printf("\tMax: %v\n", max)
	//fmt.Printf("\tAverage: %v\n", average)
}

func foo() {

}
