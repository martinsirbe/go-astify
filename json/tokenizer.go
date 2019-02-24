package json

import (
	"fmt"
	"unicode"
)

const (
	// UNKNOWN used when didn't match any of expected token types
	UNKNOWN TokenType = iota
	// EOF identifies the end of the input
	EOF
	// COLON identifies a colon separator
	COLON
	// COMMA identifies a comma separator
	COMMA
	// NULL identifies a null literal
	NULL
	// STRING identifies a string literal
	STRING
	// INTEGER identifies an integer literal
	INTEGER
	// BOOLEAN identifies a boolean literal
	BOOLEAN
	// LCBRACKET identifies a left curly bracket separator
	LCBRACKET
	// RCBRACKET identifies a right curly bracket separator
	RCBRACKET
	// LSBRACKET identifies a left square bracket separator
	LSBRACKET
	// RSBRACKET identifies a right square bracket separator
	RSBRACKET

	colonTokenValue              = ':'
	commaTokenValue              = ','
	leftCurlyBracketTokenValue   = '{'
	rightCurlyBracketTokenValue  = '}'
	leftSquareBracketTokenValue  = '['
	rightSquareBracketTokenValue = ']'
	eofTokenValue                = "eof"
)

var (
	// these chars will be skipped during the JSON object scan
	skipChars         = []rune{'\n', '\t', '\r', ' '}
	boolTokenValue    string
	integerTokenValue string
	stringTokenValue  string
)

// TokenType defines the type of the token based on JSON syntax
type TokenType int

// String provides human readable token type string
func (tt TokenType) String() string {
	return []string{
		"unknown",
		"eof",
		"colon",
		"comma",
		"null",
		"string",
		"integer",
		"boolean",
		"left curly bracket",
		"right curly bracket",
		"left square bracket",
		"right square bracket",
	}[tt]
}

// Token lexically represents a part of the JSON syntax structure
type Token struct {
	Value string
	Type  TokenType
}

// Tokenizer parses a JSON object and builds a lexical representation of it
type Tokenizer struct {
	charsIn     []rune
	currentChar rune
	start, end  int
}

// NewTokenizer initialises a new instance of the tokenizer
func NewTokenizer(charsIn []rune) *Tokenizer {
	scanner := Tokenizer{charsIn: charsIn}
	scanner.scanNextChar()
	return &scanner
}

// PeekToken preview the next token from the char input
func (t *Tokenizer) PeekToken() *Token {
	start := t.start
	end := t.end
	tok := t.GetToken()

	t.start = start
	t.end = end
	t.currentChar = t.charsIn[t.end-1]

	return tok
}

// GetToken get the next token from the char input
func (t *Tokenizer) GetToken() *Token {
	for {
		if !shouldSkipChar(t.currentChar) {
			break
		}
		t.scanNextChar()
	}

	token := t.getToken()

	t.scanNextChar()
	return &token
}

func (t *Tokenizer) getToken() Token {
	switch {
	case t.currentChar == colonTokenValue:
		return Token{Type: COLON, Value: string(colonTokenValue)}
	case t.currentChar == commaTokenValue:
		return Token{Type: COMMA, Value: string(commaTokenValue)}
	case t.currentChar == leftCurlyBracketTokenValue:
		return Token{Type: LCBRACKET, Value: string(leftCurlyBracketTokenValue)}
	case t.currentChar == rightCurlyBracketTokenValue:
		return Token{Type: RCBRACKET, Value: string(rightCurlyBracketTokenValue)}
	case t.currentChar == leftSquareBracketTokenValue:
		return Token{Type: LSBRACKET, Value: string(leftSquareBracketTokenValue)}
	case t.currentChar == rightSquareBracketTokenValue:
		return Token{Type: RSBRACKET, Value: string(rightSquareBracketTokenValue)}
	case t.isNull():
		return Token{Type: NULL, Value: ""}
	case t.isString():
		return Token{Type: STRING, Value: stringTokenValue}
	case t.isInteger():
		return Token{Type: INTEGER, Value: integerTokenValue}
	case t.isBoolean():
		return Token{Type: BOOLEAN, Value: boolTokenValue}
	case t.currentChar == rune(0):
		return Token{Type: EOF, Value: eofTokenValue}
	}
	return Token{Type: UNKNOWN, Value: string(t.currentChar)}
}

func (t *Tokenizer) scanNextChar() {
	t.currentChar = 0
	if !t.isLastChar() {
		t.currentChar = t.charsIn[t.end]
	}
	t.start = t.end
	t.end++
}

func shouldSkipChar(currentChar rune) bool {
	for _, skipChar := range skipChars {
		if currentChar == skipChar {
			return true
		}
	}
	return false
}

func (t *Tokenizer) getBooleanValue() string {
	switch {
	case t.currentChar == 't':
		return string(t.charsIn[t.start : t.start+4])
	case t.currentChar == 'f':
		return string(t.charsIn[t.start : t.start+5])
	}
	return ""
}

func (t *Tokenizer) isLastChar() bool {
	return t.end >= len(t.charsIn)
}

func (t *Tokenizer) isNull() bool {
	if t.currentChar != 'n' && t.start+4 <= len(t.charsIn) {
		return false
	}

	if string(t.charsIn[t.start:t.start+4]) == "null" {
		t.end += 4
		t.currentChar = t.charsIn[t.end]
		return true
	}
	return false
}

func (t *Tokenizer) isBoolean() bool {
	if t.currentChar != 't' && t.currentChar != 'f' && t.start+5 <= len(t.charsIn) {
		return false
	}

	bv := t.getBooleanValue()
	switch {
	case bv != "true" && bv != "false":
		return false
	case bv == "true":
		t.end += 4
	case bv == "false":
		t.end += 5
	}
	t.currentChar = t.charsIn[t.end]
	boolTokenValue = bv
	return true
}

func (t *Tokenizer) isString() bool {
	if t.currentChar != '"' {
		return false
	}

	stringTokenValue = fmt.Sprintf("%c", t.charsIn[t.end])

	for i := t.end + 1; i < len(t.charsIn); i++ {
		if t.charsIn[i] != '"' {
			stringTokenValue = fmt.Sprintf("%s%c", stringTokenValue, t.charsIn[i])
			continue
		}
		t.end = i + 1
		t.currentChar = t.charsIn[t.end]
		return true
	}
	return false
}

func (t *Tokenizer) isInteger() bool {
	if t.currentChar != '-' && !unicode.IsDigit(t.currentChar) {
		return false
	}

	integerTokenValue = fmt.Sprintf("%c%c", t.currentChar, t.charsIn[t.end])

	t.end++
	t.currentChar = t.charsIn[t.end]
	for unicode.IsDigit(t.currentChar) {
		integerTokenValue = fmt.Sprintf("%s%c", integerTokenValue, t.charsIn[t.end])
		t.end++
		t.currentChar = t.charsIn[t.end]
	}
	return true
}
