package stringUtil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string, b *strings.Builder) {
	isASCII, hasUpper := true, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			break
		}
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}

	if isASCII { // optimize for ASCII-only strings.
		if !hasUpper {
			b.WriteString(s)
			return
		}
		var (
			pos int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'A' <= c && c <= 'Z' {
				c += 'a' - 'A'
				if pos < i {
					b.WriteString(s[pos:i])
				}
				b.WriteByte(c)
				pos = i + 1
			}
		}
		if pos < len(s) {
			b.WriteString(s[pos:])
		}
		return
	}
	Map(unicode.ToLower, s, b)
	return
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string, b *strings.Builder) {
	isASCII, hasLower := true, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			break
		}
		hasLower = hasLower || ('a' <= c && c <= 'z')
	}

	if isASCII { // optimize for ASCII-only strings.
		if !hasLower {
			b.WriteString(s)
			return
		}
		var (
			pos int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
				if pos < i {
					b.WriteString(s[pos:i])
				}
				b.WriteByte(c)
				pos = i + 1
			}
		}
		if pos < len(s) {
			b.WriteString(s[pos:])
		}
		return
	}
	Map(unicode.ToUpper, s, b)
	return
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string, b *strings.Builder) {
	// In the worst case, the string can grow when mapped, making
	// things unpleasant. But it's so rare we barge in assuming it's
	// fine. It could also shrink but that falls out naturally.

	// The output buffer b is initialized on demand, the first
	// time a character differs.
	for i, c := range s {
		r := mapping(c)
		if r == c && c != utf8.RuneError {
			continue
		}

		var width int
		if c == utf8.RuneError {
			c, width = utf8.DecodeRuneInString(s[i:])
			if width != 1 && r == c {
				continue
			}
		} else {
			width = utf8.RuneLen(c)
		}

		b.Grow(len(s) + utf8.UTFMax)
		b.WriteString(s[:i])
		if r >= 0 {
			b.WriteRune(r)
		}

		s = s[i+width:]
		break
	}

	// Fast path for unchanged input
	if b.Cap() == 0 { // didn't call b.Grow above
		b.WriteString(s)
		return
	}

	for _, c := range s {
		r := mapping(c)

		if r >= 0 {
			// common case
			// Due to inlining, it is more performant to determine if WriteByte should be
			// invoked rather than always call WriteRune
			if r < utf8.RuneSelf {
				b.WriteByte(byte(r))
			} else {
				// r is not an ASCII rune.
				b.WriteRune(r)
			}
		}
	}
}
