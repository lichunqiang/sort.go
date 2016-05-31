package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fixture := []int{4, 18, 3, 89, 12, 45}

	Bubble(fixture)
	fmt.Println("Bubble sort result:")
	fmt.Println(fixture)
}
