/*
{
	"name": "Daniel Hernandez",
	"birthDate": "18/03/1988",
	"salaries": [
		{
			date: "01/01/2018"
			salary: 100.00,
			bonus: 10.10
		},
		{
			date: "01/01/2018"
			salary: 200.00,
			bonus: 10.10
		}
	],
	"stats": {
		"average": 100,
		"min": 110.10,
		"max": 210.10
	}
}
*/

package main

import (
	"employee"
	"encoding/json"
	"io/ioutil"
)

// cd src\employeedata\generator && go install
func main() {
	employees := employee.GenerateRandomData(100)
	employeeBytes, _ := json.Marshal(employees)
	ioutil.WriteFile("outputsfile/employees-json.json", employeeBytes, 0644)
}
