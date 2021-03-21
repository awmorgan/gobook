package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty string", "", true},
		{"single letter", "a", true},
		{"two same letters", "aa", true},
		{"two different letters", "ab", false},
		{"word1", "kayak", true},
		{"word2", "detartrated", true},
		{"sentence1", "A man, a plan, a canal: Panama", true},
		{"sentence2", "Evil I did dwell; lewd did I live.", true},
		{"sentence3", "Able was I ere I saw Elba", true},
		{"French character", "été", true},
		{"Foreign language", "Et se resservir, ivresse reste.", true},
		{"Non-palindrome", "palindrome", false},
		{"Semi-palindrome", "desserts", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
