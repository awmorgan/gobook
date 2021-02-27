package main

import "fmt"

func main() {
	s := []string{"", "a", "", "b", "c", "", "", "d"}
	s1 := s
	fmt.Println("s: ", s)
	fmt.Println("nonempty(s): ", nonempty(s))
	fmt.Println("s: ", s)

	fmt.Println("s1: ", s1)
	s1 = nonempty2(s1)
	fmt.Println("s1 = nonempty2(s1); s1: ", s1)

}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		} c
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
