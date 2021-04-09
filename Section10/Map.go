package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS IF STATEMENT", v)
	}
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS IF STATEMENT", v)
	}

	// add element & range
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete
	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)
}
