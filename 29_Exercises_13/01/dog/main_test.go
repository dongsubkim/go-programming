package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}
func BenchmarkYears(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}
func ExampleYearsTwo() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("Expected", 70, "got", n)
	}
}
func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Expected", 70, "got", n)
	}
}
