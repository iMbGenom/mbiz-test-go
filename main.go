package main

import (
	"fmt"
	"strconv"
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
	number := 5
	case3 := getFactorial(number)
	fmt.Println(case3)

	// case #4
	case4_1 := numeralIntoWords("5:01") // with parameter: string
	case4_2 := _getWords(5, 28)         // with parameters: hour and minute
	case4_3 := _getWords(5, 60)         // wrong time format
	fmt.Println(case4_1)
	fmt.Println(case4_2)
	fmt.Println(case4_3)
}

func numeralIntoWords(time string) string {
	hour, minute := _sanitizeTime(time)
	return _getWords(hour, minute)
}

func _sanitizeTime(time string) (int, int) {
	timeArr := strings.Split(time, ":")
	hour, err := strconv.Atoi(timeArr[0])
	minute, err := strconv.Atoi(timeArr[1])
	if err != nil {
		return 0, 0
	}
	return hour, minute
}

func _getWords(hour, minute int) string {
	fmt.Println(hour, minute)
	var result, spell string
	validMinute := true
	finalHour := _numberToWord(hour)
	finalMinute := _numberToWord(minute)
	if minute >= 60 {
		validMinute = false
		return "Wrong Time Format case #4"
	}

	if validMinute && minute == 0 {
		spell = " O'clock"
		result = finalHour + spell
	}
	if validMinute && minute == 15 {
		spell = "Quarter Past "
		result = spell + finalHour
	}
	if validMinute && minute == 30 {
		spell = "Half Past "
		result = spell + finalHour
	}
	if validMinute && minute == 45 {
		spell = "Quarter To "
		finalHour = _numberToWord(hour + 1)
		result = spell + finalHour
	}
	if validMinute && minute < 30 && minute != 0 && minute != 15 {
		spell = " Minutes Past "
		result = finalMinute + spell + finalHour
	}
	if validMinute && minute > 30 && minute != 45 {
		spell = " Minutes To "
		finalMinute := _numberToWord(60 - minute)
		result = finalMinute + spell + finalHour
	}

	return result
}

func _numberToWord(number int) string {
	var result string
	switch number {
	case 1:
		result = "One"
	case 2:
		result = "Two"
	case 3:
		result = "Three"
	case 4:
		result = "Four"
	case 5:
		result = "Five"
	case 6:
		result = "Six"
	case 7:
		result = "Seven"
	case 8:
		result = "Eight"
	case 9:
		result = "Nine"
	case 10:
		result = "Ten"
	case 13:
		result = "Thirteen"
	case 20:
		result = "Twenty"
	case 28:
		result = "Twenty Eight"
	default:
		result = "Undefined"
	}
	return result
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
