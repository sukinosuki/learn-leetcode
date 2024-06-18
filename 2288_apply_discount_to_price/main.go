package main

import (
	"fmt"
	"strconv"
)

func discountPrices(sentence string, discount int) string {

	var ans []byte

	i := 0

	n := len(sentence)
	discountF := float64(discount) / 100
	for i < n {

		//if (sentence[i] == '$' && i < n-1) || i == 0 && sentence[i-1] == ' ' {
		if (i == 0 && sentence[i] == '$') || (i > 0 && sentence[i-1] == ' ' && sentence[i] == '$') {
			ans = append(ans, '$')

			j := i + 1

			for j < n && isNumeric(sentence[j]) {
				j++
			}
			// 到达最后并且最后一个字符不是空格
			if j == n {
				if !isNumeric(sentence[j-1]) && sentence[j-1] != ' ' {

					i++
					continue
				}
			} else {
				if !isNumeric(sentence[j]) && sentence[j] != ' ' {

					i++
					continue
				}
			}

			if j > i+1 {
				num, _ := strconv.Atoi(sentence[i+1 : j])

				sub := float64(num) - float64(num)*discountF

				s := fmt.Sprintf("%.2f", sub)
				ans = append(ans, []byte(s)...)

				i = j - 1
			}
		} else {
			ans = append(ans, sentence[i])
		}

		i++
	}

	return string(ans)
}

func isNumeric(s byte) bool {
	switch s {
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		return true
	}

	return false
}

func main() {
	// there are $0.50 $1.00 and 5$ candies in the shop
	//sentense := "there are $1 $2 and 5$ candies in the shop"
	//discount := 50

	// "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
	// 1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $0.00$

	//sentense := "1 2 $3 4 $5 $6 7 8$ $9 $10$"
	//discount := 100
	// $2$3 $10.00 $100.00 $1.00 200 $33.00 33$ $$ $99.00 $99999.00 $9999999999.00
	sentense := "$2$3 $10 $100 $1 200 $33 33$ $$ $99 $99999 $9999999999"
	discount := 0

	//  ka3caz4837h6ada4 r1 $547.82
	//sentense := "ka3caz4837h6ada4 r1 $602"
	//discount := 9

	res := discountPrices(sentense, discount)

	fmt.Printf("res: %v\n", res)
}
