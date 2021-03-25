package word

import (
	"math/rand"
	"testing"
	"time"
)

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

// randomPalindrome returns a palindrome whose length and contents
// are derived from teh pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
