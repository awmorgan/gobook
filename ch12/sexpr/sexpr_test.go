package sexpr_test

import (
	"testing"

	"github.com/awmorgan/gobook/ch12/sexpr"
	"github.com/google/go-cmp/cmp"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	// Color           bool
	Actor  map[string]string
	Oscars []string
	Sequel *string
}

func TestMarshalUnmarshal(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		// Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	b, err := sexpr.Marshal(strangelove)
	if err != nil {
		t.Errorf("Marshal failed: %v", err)
	}
	var m Movie
	if err := sexpr.Unmarshal(b, &m); err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}
	if !cmp.Equal(m, strangelove) {
		t.Errorf("m != strangelove: %s", cmp.Diff(m, strangelove))
	}
}
