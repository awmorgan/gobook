package main

import "fmt"

func main() {
	s := []string{"", "a", "", "b", "c", "", "", "d"}
	fmt.Println(s)
	fmt.Println(nonempty(s))
}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
