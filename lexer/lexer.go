package lexer

import (
	"fmt"
	"strings"
)

type lexer struct {
	tokens []Token
	source string
	pos    int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) atAdvanceN(n int) byte {
	return lex.source[lex.pos+n]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		tokens: make([]Token, 0),
	}
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isLeftParenthese(c byte) bool {
	return c == '('
}

func isRightParenthese(c byte) bool {
	return c == ')'
}

func isLeftBrace(c byte) bool {
	return c == '{'
}

func isRightBrace(c byte) bool {
	return c == '}'
}

func isLeftBracket(c byte) bool {
	return c == '['
}

func isRightBracket(c byte) bool {
	return c == ']'
}

func isColon(c byte) bool {
	return c == ':'
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	// loop when there are still tokens remaining
	for !lex.at_eof() {
		if isWhitespace(lex.at()) {

		} else if isAlphaNumeric(lex.at()) {
			var builder strings.Builder

			for isAlphaNumeric(lex.at()) && !lex.at_eof() {
				builder.WriteString(string(lex.at()))
				lex.advanceN(1)
			}

			value := builder.String()

			if kind, exists := reservedLookup[value]; exists {
				lex.push(newToken(value, kind))
			} else {
				lex.push(newToken(value, IDENTIFIER))
			}
			//skip the advance at the end
			continue

		} else if isLeftParenthese(lex.at()) {
			lex.push(newToken("", LPAREN))
		} else if isRightParenthese(lex.at()) {
			lex.push(newToken("", RPAREN))
		} else if isLeftBrace(lex.at()) {
			lex.push(newToken("", LBRACE))
		} else if isRightBrace(lex.at()) {
			lex.push(newToken("", RBRACE))
		} else if isLeftBracket(lex.at()) {
			lex.push(newToken("", LBRACKET))
		} else if isRightBracket(lex.at()) {

			lex.push(newToken("", RBRACKET))
		} else if isColon(lex.at()) {
			lex.push(newToken("", COLON))
		} else if lex.at() == '/' && lex.atAdvanceN(1) == '/' {
			//comments
			lex.advanceN(2)
			var builder strings.Builder

			for lex.at() != '\n' && !lex.at_eof() {
				builder.WriteString(string(lex.at()))
				lex.advanceN(1)
			}

			lex.push(newToken(builder.String(), SINGLE_LINE_COMMENT))
		} else if kind, exists := operatorLookup[string(lex.at())]; exists {
			lex.push(newToken(string(lex.at()), kind))
		} else {
			fmt.Printf("Unknown character: %c\n", lex.at())
		}

		lex.advanceN(1)
	}

	lex.tokens = append(lex.tokens, newToken("EOF", EOF))
	return lex.tokens
}

// potential bug
// o and x as operators could also be identifieres as o or x
// handle later
