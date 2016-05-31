package sort

import (
	"fmt"
)

//最好的时间复杂度是O(n)
//最坏的时间复杂度是O(n2)
func Bubble(values []int) {
	count := len(values)

	for i := 0; i < count; i++ {
		for j := 0; j < count-i-1; j++ {
			if values[j] > values[j+1] {
				values[j+1], values[j] = values[j], values[j+1]
			}
			fmt.Println(values)
		}
	}
}
