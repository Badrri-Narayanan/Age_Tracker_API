package utils

import (
	"time"
)

func IsLeapYear(year int) bool {
	// years which are divisible by 4 except centuries are leap year. If a century year is divisible by 400, then it's a leap year
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func GetDateDifference(startDate, endDate time.Time) (int, int, int) {
	day1 := startDate.Day()
	month1 := startDate.Month()
	year1 := startDate.Year()

	day2 := endDate.Day()
	month2 := endDate.Month()
	year2 := endDate.Year()

	// if end day is less than start day, we need to carry-forward some days from end month
	if day2 < day1 {
		// since Feb has only 28 or 29 days
		if month2 == 3 {
			// leap years has 29 days in Feb
			if IsLeapYear(year2) {
				day2 += 29
			} else {
				day2 += 28
			}
		} else if month2 == 5 || month2 == 7 || month2 == 10 || month2 == 12 {
			// pervious months of May, July, October & Decemeber which are April, June, Sept & Nov have 30 days
			day2 += 30
		} else {
			// other pervious months all have 31 days
			day2 += 31
		}
		// we have carry-forwarded previous month to be used as days
		month2--
	}

	// if end month is less then start month, we need to carry-forward from end year
	if month2 < month1 {
		month2 += 12
		year2--
	}

	dayDiff := day2 - day1
	monthDiff := int(month2) - int(month1)
	yearDiff := year2 - year1

	return dayDiff, monthDiff, yearDiff
}

func CalculateAgeInMonthAndYears(dateOfBirth time.Time) (int, int, int) {
	today := time.Now()

	return GetDateDifference(dateOfBirth, today)
}
