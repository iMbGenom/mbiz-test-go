package main

import (
	"fmt"
	"strings"
)

func main() {
	// case #1
	arr := []int{1, 2, 3, 4}
	case1 := sumArrayValues(arr)
	fmt.Println(case1)

	// case #2
	fullDate := "Tue 16 Jul 2019"
	case2 := changeDateFormat(fullDate)
	fmt.Println(case2)

	// case #3
	case3 := getFactorial(5)
	fmt.Println(case3)

	// case #4

}

func getFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * getFactorial(num-1)
}

func changeDateFormat(fullDate string) string {
	dateArr := strings.Split(fullDate, " ")

	if len(dateArr) != 4 || len(dateArr[0]) != 3 || len(dateArr[2]) != 3 || len(dateArr[3]) != 4 {
		return "Wrong Date Format in case #2"
	}

	dayDate, monthName, year := dateArr[1], dateArr[2], dateArr[3]
	validMonth := _changeMonthName(monthName)

	return year + "-" + validMonth + "-" + dayDate
}

func _changeMonthName(monthName string) string {
	var result string
	switch monthName {
	case "Jan":
		result = "01"
	case "Feb":
		result = "02"
	case "Mar":
		result = "03"
	case "Apr":
		result = "04"
	case "May":
		result = "05"
	case "Jun":
		result = "06"
	case "Jul":
		result = "07"
	case "Aug":
		result = "08"
	case "Sep":
		result = "09"
	case "Oct":
		result = "10"
	case "Nov":
		result = "11"
	case "Dec":
		result = "12"
	default:
		result = "undefined"
	}
	return result
}

func sumArrayValues(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
