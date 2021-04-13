package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["john_doe"] = []string{"abc", "def", "ghi"}
	for k, v := range m {
		fmt.Printf("Key %s, record: %v\n", k, v)
	}
	delete(m, "john_doe")
	for k, v := range m {
		fmt.Printf("Key %s, record: %v\n", k, v)
	}
}
