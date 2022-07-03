package model

import (
	"main/utils"
	"time"
)

type People struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Age         Age       `json:"age"`
}

func (people *People) CalculateAge() {
	people.Age.Year, people.Age.Month = utils.CalculateAgeInMonthAndYears(&people.DateOfBirth)
}
