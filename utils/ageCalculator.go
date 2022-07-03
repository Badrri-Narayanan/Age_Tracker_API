package utils

import (
	"time"
)

func getDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func CalculateAgeInMonthAndYears(dateOfBirth *time.Time) (int, int) {
	today := time.Now()

	var monthDiff int
	yearDiff := getDiff(today.Year(), dateOfBirth.Year())

	if int(today.Month()) < int(dateOfBirth.Month()) {
		monthDiff = int(dateOfBirth.Month()) - 1
		yearDiff--
	} else {
		monthDiff = int(today.Month()) - int(dateOfBirth.Month())
	}

	return yearDiff, monthDiff
}
