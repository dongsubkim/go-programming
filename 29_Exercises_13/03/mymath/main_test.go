package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		numbers []int
		answer  float64
	}

	tests := []test{
		test{[]int{1, 2, 3}, 2},
		test{[]int{1, 2, 3, 4}, 2.5},
	}

	for _, v := range tests {
		a := CenteredAvg(v.numbers)
		if a != v.answer {
			t.Error("Expected", v.answer, "Got", a)
		}
	}

}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4}))
	// Output:
	// 2.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}
