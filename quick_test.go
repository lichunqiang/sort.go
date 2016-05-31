package sort

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	fixture := []int{4, 18, 3, 89, 12, 45}

	Quick(fixture)
	fmt.Println("Quick sort result:")
	fmt.Println(fixture)
}
