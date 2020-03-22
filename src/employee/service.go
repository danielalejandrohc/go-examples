package employee

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomData(n int) []Employee {
	var salaries []Salary
	var employees []Employee
	layout := "2006-01-02"
	for j := 0; j < n; j++ {
		employeeRecord := Employee{
			Name:      "Daniel",
			BirthDate: time.Now(),
		}
		for i := 1; i <= 12; i++ {
			twoDigitsFormat := fmt.Sprintf("%02d", i)
			dateAsString := "2020-" + twoDigitsFormat + "-01"
			dateAsTime, _ := time.Parse(layout, dateAsString)
			x := Salary{
				Date:   dateAsTime,
				Salary: float32(100 + (i * i)),
			}
			x.Bonus = 100 * rand.Float32()
			salaries = append(salaries, x)
		}
		employeeRecord.Salaries = salaries
		employees = append(employees, employeeRecord)
		fmt.Printf("%v ", j)
	}
	return employees
}
