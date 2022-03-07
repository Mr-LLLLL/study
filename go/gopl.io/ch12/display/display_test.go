package display

import (
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestDisplay(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How i learned to stop worrying and love the bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "peter sellers",
			"Grp.Capt. lionel mandrake":  "peter sellers",
			"pres. merkin mufley":        "peter sellers",
			"gen. bck turgidson":         "george c. scott",
			"brig. gen. gack d. rippeer": "steling hatden",
			`maj. t.j. "king" long`:      "slim picknens",
		},
		Oscars: []string{
			"best actor (nomin.)",
			"best adapterd screenplay (nomin.)",
			"best director (nomin.)",
			"vest picture (nomin.)",
		},
	}

	Display("strangelove", strangelove)
}

func TestTest(t *testing.T) {
	// Display("os.Stderr", os.Stderr)
	// Display("rV", reflect.ValueOf(os.Stderr))
	var i interface{} = 3

	// Display("i", i)
	Display("&i", &i)
}
