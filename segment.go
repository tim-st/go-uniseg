// Package uniseg implements basic unicode segmentation.
package uniseg

import (
	"fmt"
	"unicode/utf8"

	"github.com/tim-st/go-unicat"
)

type countType = uint8

// Segment is a part of a given byte slice with the property that
// the segment is either a word or each rune in the segment is in the same
// General Unicode Category.
type Segment struct {
	Segment   []byte
	Category  Category
	RuneCount countType
}

func (s Segment) String() string {
	return fmt.Sprintf(`Segment{"%s" (Category: %s; Runes: %d)}`, s.Segment, s.Category, s.RuneCount)
}

// Segments is a Slice of Segments of the given text.
func Segments(text []byte) []Segment {
	if len(text) == 0 {
		return nil
	}
	var segments = make([]Segment, 0, len(text)/3) // TODO: cap ok?

	firstRuneIsLetterUppercase := false
	numberLowercaseLettersInSegment := countType(0)
	numberUppercaseLettersInSegment := countType(0)
	segmentStartIdx := 0
	r, i := utf8.DecodeRune(text[:])
	runeCount := countType(0)
	prevCat := unicat.From(r)
	for {
		runeCount++
		switch prevCat {
		case unicat.UnicodeLl:
			numberLowercaseLettersInSegment++
		case unicat.UnicodeLu:
			if runeCount == 1 {
				firstRuneIsLetterUppercase = true
			}
			numberUppercaseLettersInSegment++
		}

		atEnd := false
		currentCat := unicat.UnicodeCn
		oldIdx := i
		if i < len(text) {
			currentRune, rSize := utf8.DecodeRune(text[i:])
			i += rSize
			currentCat = unicat.From(currentRune)
			if prevCat.IsLetter() && currentCat.IsLetter() {
				prevCat = currentCat
				continue
			}
		} else {
			atEnd = true
		}

		if atEnd || prevCat != currentCat {

			segment := Segment{
				Segment:   text[segmentStartIdx:oldIdx],
				Category:  Category(prevCat),
				RuneCount: runeCount,
			}

			if runeCount > 1 && prevCat.IsLetter() {
				switch numberUppercaseLettersInSegment {
				case 0:
					segment.Category = WordAllLower
				case 1:
					if firstRuneIsLetterUppercase {
						segment.Category = WordFirstUpper
					} else {
						segment.Category = WordMixedLetters
					}
				case runeCount:
					segment.Category = WordAllUpper
				default:
					segment.Category = WordMixedLetters
				}
			}

			segments = append(segments, segment)

			if atEnd {
				return segments
			}

			prevCat = currentCat
			segmentStartIdx = oldIdx
			runeCount = 0
			numberLowercaseLettersInSegment = 0
			numberUppercaseLettersInSegment = 0
			firstRuneIsLetterUppercase = false
		}

	}
}
