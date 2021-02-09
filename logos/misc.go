package logos

import (
	"strings"
	"unicode"

	runewidth "github.com/mattn/go-runewidth"
)

// splits a string into lines by newline.
func splitLines(s string) (ss []string) {
	return strings.Split(s, "\n")
}

// splits a string according to unicode spaces.
func splitSpaces(s string) (ss []string) {
	buf := []rune{}
	for _, r := range s {
		if unicode.IsSpace(r) {
			// continue
			if len(buf) > 0 {
				ss = append(ss, string(buf))
				buf = nil
			}
		} else {
			buf = append(buf, r)
		}
	}
	if len(buf) > 0 {
		ss = append(ss, string(buf))
		// buf = nil
	}
	return ss
}

// gets the terminal display width of a string.
// must be compatible with nextCharacter().
func widthOf(s string) (l int) {
	zwj := false // zero width joiner '\u200d'.
	for _, r := range s {
		if r == '\u200d' {
			zwj = true
			continue
		}
		if zwj {
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			l++ // show a blank instead?
		case 1:
			l++
		case 2:
			l += 2
		default:
			panic("should not happen")
		}
	}
	return l
}

func toRunes(s string) []rune {
	runes := make([]rune, 0, len(s))
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

/*
// XXX DEPRECATED
// splits a string into two parts, returning
// the longest string of given width as first result,
// and the remaining string as second.
// if w is the width of s or greater, p2 is empty.
func splitWidth(s string, w int) (p1, p2 []rune) {
	var l int = 0
	p1rz := make([]rune, len(s))
	p2rz := make([]rune, len(s))
	zwj := false // zero width joiner '\u200d'.
	for _, r := range str {
		if l < w {
			p1rz = append(p1rz, r)
		} else {
			p2rz = append(p2rz, r)
		}
		if r == '\u200d' {
			zwj = true
			continue
		}
		if zwj {
			l++
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			l++ // show a blank instead?
		case 1:
			l++
		case 2:
			l += 2
		}
	}
	return p1rz, p2rz
}
*/

// given runes of a valid utf8 string,
// return a string that represents
// the next single character (with any modifiers).
// w: width of character. n: number of runes read
func nextCharacter(rz []rune) (s string, w int, n int) {
	for n = 0; n < len(rz); n++ {
		r := rz[n]
		if r == '\u200d' {
			// special case: zero width joins.
			s = s + string(r)
			if n+1 < len(rz) {
				s = s + string(rz[n+1])
				n++
				continue
			} else {
				// just continue, return invalid string s.
				n++
				return
			}
		} else if 0 < len(s) {
			return
		} else {
			// append r to s and inc w.
			rw := runewidth.RuneWidth(r)
			s = s + string(r)
			if rw == 0 {
				w += 1
			} else {
				w += rw
			}
		}
	}
	return
}
