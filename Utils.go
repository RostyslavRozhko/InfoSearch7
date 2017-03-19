package main

import (
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanTerms(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var start int
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isControlBreak(r) {
			break
		}
	}

	for i := start; i < len(data); {
		r, width := utf8.DecodeRune(data[i:])
		if isControlBreak(r) {
			return i + width, data[start:i], nil
		}
		i += width
	}

	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

func isPunctuation(r rune) bool {

	if r == '\'' || r == '-' {
		return false
	}
	if r >= '!' && r <= '/' ||
		r >= ':' && r <= '@' ||
		r >= '[' && r <= '`' ||
		r >= '{' && r <= '~' {
		return true
	}

	return false
}

func isControlBreak(r rune) bool {
	return isSpace(r) || isPunctuation(r)
}

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

func removePunctuation(r rune) rune {
	if strings.ContainsRune(".,:;?!'&()`\"{}|[]#$%_*/\\><[]^`", r) {
		return -1
	}
	return r
}
