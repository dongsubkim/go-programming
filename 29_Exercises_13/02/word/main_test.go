package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	s := "Shaken not stirred"
	c := Count(s)
	if c != 3 {
		t.Error("Expected", 3, "Got", c)
	}
}
func TestCountTwo(t *testing.T) {
	s := "Shaken not stirred"
	c := CountTwo(s)
	if c != 3 {
		t.Error("Expected", 3, "Got", c)
	}
}

func ExampleCount() {
	s := "Shaken not stirred"
	c := Count(s)
	fmt.Println(c)
	// Output:
	// 3
}
func ExampleCountTwo() {
	s := "Shaken not stirred"
	c := CountTwo(s)
	fmt.Println(c)
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	s := "Shaken not stirred"
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}

func BenchmarkCountTwo(b *testing.B) {
	s := "Shaken not stirred"
	for i := 0; i < b.N; i++ {
		CountTwo(s)
	}
}
