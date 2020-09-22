package calculate

import (
	"errors"
	"fmt"
)

//SumX returns the sum of all the numbers passed as parameters to the function
// func SumX(numbers ...int) int {
// 	sum := 0
// 	for _, num := range numbers {
// 		sum = sum + num
// 	}
// 	return sum
// }

//SumX returns the sum of all the numbers passed as parameters to the function
func SumX(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("provide 2 or more")
	} else {
		for _, x := range numbers {
			sum = sum + x
		}
		return sum, nil
	}

}

//Sum computes the sum of two numbers
func Sum(x, y int) int {
	return x + y
}

//Fibonnaci computes the nth fibonnaci series
func Fibonnaci(n int) string {
	if n == 1 {
		return "0"
	} else {
		res := "0 1 "
		i := 0
		j := 1
		for x := 3; x <= n; x++ {
			temp := i + j
			i = j
			j = temp
			res = res + fmt.Sprint(temp) + " "
		}
		return res
	}

}
