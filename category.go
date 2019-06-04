package uniseg

import (
	"github.com/tim-st/go-unicat"
)

// Category is the Category of a segment.
// It's identified by the General Unicode Category and semantical analysis of the rune sequence.
// The base type is unicat.Category which has semantics for single runes, while this
// type extends it to better describe sequences of runes like words.
type Category uint8

const (
	// UnicodeLu = Letter, uppercase
	UnicodeLu Category = iota
	// UnicodeLl = Letter, lowercase
	UnicodeLl
	// UnicodeLt = Letter, titlecase
	UnicodeLt
	// UnicodeLm = Letter, modifier
	UnicodeLm
	// UnicodeLo = Letter, other
	UnicodeLo
	// UnicodeMn = Mark, nonspacing
	UnicodeMn
	// UnicodeMc = Mark, spacing combining
	UnicodeMc
	// UnicodeMe = Mark, enclosing
	UnicodeMe
	// UnicodeNd = Number, decimal digit
	UnicodeNd
	// UnicodeNl = Number, letter
	UnicodeNl
	// UnicodeNo = Number, other
	UnicodeNo
	// UnicodePc = Punctuation, connector
	UnicodePc
	// UnicodePd = Punctuation, dash
	UnicodePd
	// UnicodePs = Punctuation, open
	UnicodePs
	// UnicodePe = Punctuation, close
	UnicodePe
	// UnicodePi = Punctuation, initial quote
	UnicodePi
	// UnicodePf = Punctuation, final quote
	UnicodePf
	// UnicodePo = Punctuation, other
	UnicodePo
	// UnicodeSm = Symbol, math
	UnicodeSm
	// UnicodeSc = Symbol, currency
	UnicodeSc
	// UnicodeSk = Symbol, modifier
	UnicodeSk
	// UnicodeSo = Symbol, other
	UnicodeSo
	// UnicodeZs = Separator, space
	UnicodeZs
	// UnicodeZl = Separator, line
	UnicodeZl
	// UnicodeZp = Separator, paragraph
	UnicodeZp
	// UnicodeCc = Other, control
	UnicodeCc
	// UnicodeCf = Other, format
	UnicodeCf
	// UnicodeCs = Other, surrogate
	UnicodeCs
	// UnicodeCo = Other, private use
	UnicodeCo
	// UnicodeCn = Other, not assigned
	UnicodeCn
	// WordAllLower = UnicodeLl{2,}
	WordAllLower
	// WordFirstUpper = UnicodeLu{1} + UnicodeLl{1,}
	WordFirstUpper
	// WordAllUpper = UnicodeLu{2,}
	WordAllUpper
	// WordMixedLetters = aB, abC, AbC, aBc, aBC, ...
	WordMixedLetters
)

func (c Category) String() (s string) {
	switch c {
	case WordAllLower:
		s = "Wl"
	case WordFirstUpper:
		s = "Wf"
	case WordAllUpper:
		s = "Wu"
	case WordMixedLetters:
		s = "Wm"
	default:
		s = unicat.Category(c).String()
	}
	return
}

// IsLetter checks if the category is Letter (L).
func (c Category) IsLetter() bool { return unicat.Category(c).IsLetter() }

// IsMark checks if the category is Mark (M).
func (c Category) IsMark() bool { return unicat.Category(c).IsMark() }

// IsNumber checks if the category is Number (N).
func (c Category) IsNumber() bool { return unicat.Category(c).IsNumber() }

// IsPunctuation checks if the category is Punctuation (P).
func (c Category) IsPunctuation() bool { return unicat.Category(c).IsPunctuation() }

// IsSymbol checks if the category is Symbol (S).
func (c Category) IsSymbol() bool { return unicat.Category(c).IsSymbol() }

// IsSeparator checks if the category is Separator (Z).
func (c Category) IsSeparator() bool { return unicat.Category(c).IsSeparator() }

// IsOther checks if the category is Other (C).
func (c Category) IsOther() bool { return unicat.Category(c).IsOther() }

// IsWord checks if the Category is Word (two or more letters).
func (c Category) IsWord() bool {
	switch c {
	case WordAllLower, WordAllUpper, WordFirstUpper, WordMixedLetters:
		return true
	default:
		return false
	}
}
