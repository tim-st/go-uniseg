package uniseg

import (
	"testing"
	"unicode/utf8"
)

func TestSegmentation(t *testing.T) {

	const text = `Some test code. 10 Cafés, GitHub.`

	segments := Segments([]byte(text))

	for _, segment := range segments {
		if len(segment.Segment) == 0 {
			t.Errorf("Got an empty segment: %s", segment)
		}

		if segment.RuneCount == 0 {
			t.Errorf("RuneCount was 0: %s", segment)
		}

		if expectedLen := utf8.RuneCount(segment.Segment); expectedLen != int(segment.RuneCount) {
			t.Errorf("Expected RuneCount was %d; got: %d", expectedLen, segment.RuneCount)
		}

		if segment.Category == UnicodeCn {
			t.Errorf("Segment has unknown Category: %s", segment)
		}

	}

	const expectedLen = 14
	if l := len(segments); l != expectedLen {
		t.Errorf("Expected length of Segments was %d; got %d", expectedLen, l)
	}

}
