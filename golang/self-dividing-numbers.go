/**
 * Link of the problem
 * https://leetcode.com/problems/self-dividing-numbers/
 */

package leetcode

import (
	"fmt"
	"math"
)

func getNumber(num, pos int) int {
	r := num % int(math.Pow(10, float64(pos)))
	return r / int(math.Pow(10, float64(pos-1)))
}

func selfDividingNumbers(left int, right int) []int {
	list := make([]int, 0)
	for left <= right {
		count := 0
		length := len(fmt.Sprintf("%d", left))
		for i := length; i > 0; i-- {
			number := getNumber(left, i)
			if number != 0 && left%number == 0 {
				count++
			}
		}

		if count == length {
			list = append(list, left)
		}

		left++
	}

	return list
}
