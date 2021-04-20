// Package word provides methods to check word count in a string
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
// UseCount counts word count for each word in string s
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words in string s
func Count(s string) int {
	// write the code for this func
	return len(strings.Fields(s))
}

// CountTwo counts the number of words in string s
func CountTwo(s string) int {
	c := 0
	m := UseCount(s)
	for _, v := range m {
		c += v
	}
	return c
}
