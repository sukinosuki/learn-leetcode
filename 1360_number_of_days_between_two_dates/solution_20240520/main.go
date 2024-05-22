package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {

	strs1 := strings.Split(date1, "-")
	strs2 := strings.Split(date2, "-")

	arr1 := make([]int, 3)
	arr2 := make([]int, 3)
	for i := 0; i < 3; i++ {
		num1, _ := strconv.Atoi(strs1[i])
		num2, _ := strconv.Atoi(strs2[i])

		arr1[i] = num1
		arr2[i] = num2
	}

	return int(math.Abs(float64(dateToInt(arr1[0], arr1[1], arr1[2])) - float64(dateToInt(arr2[0], arr2[1], arr2[2]))))
}

func dateToInt(year, month, day int) int {
	monthDays := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// 1971-01-01 与 1971-01-02相差一天，所以要-1，不-1的话就是相差2天
	ans := day - 1

	// month为4时，需要加上1，2，3这三个月每月的天数
	for month > 0 {
		// 先--再获取该月的天数。4获取3月，3获取2月，2获取1月，1时获取索引0的天数(所以需要在前面多加一个0天的月份)
		month--
		ans += monthDays[month]
		if month == 2 && leapYear(year) {
			ans += 1
		}
	}
	// 与1971年相差年 * 365(还要再加上闰年的天数)
	ans += 365 * (year - 1971)

	//1. 先加上所有模 4 为 0 的年份的数量。此时有些模 100 为 0 的不是闰年的年份被加上了。
	//2. 再减去所有模 100 为 0 的年份的数量。此时有些模 400 为 0 的是闰年的年份被减去了。
	//3. 再加上所有模 400 为 0 的年份的数量。完成。
	ans += (year-1)/4 + 1971/4
	ans -= (year-1)/100 - 1971/100
	ans += (year-1)/400 - 1971/400

	return ans
}

func leapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func main() {
	//1
	//date1 := "2019-06-29"
	//date2 := "2019-06-30"

	// 15
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	res := daysBetweenDates(date1, date2)

	fmt.Printf("res: %v\n", res)
}
