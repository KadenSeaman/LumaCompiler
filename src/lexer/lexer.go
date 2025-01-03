package lexer

import (
	"fmt"
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

func isOperator(c byte) bool {
	return c == '<' || c == '>' || c == '|' || c == 'o' || c == '-' || c == '.' || c == '*' || c == 'x' || c == '#' || c == '~'
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
			str := ""
			for isAlphaNumeric(lex.at()) {
				str += string(lex.at())
				lex.advanceN(1)
			}

			if isReserved(str) {
				lex.push(newToken(str, getReservedTypeFromStr((str))))
			} else if isOperator(lex.at()) {
				lex.push(newToken(str, OPERATOR))
			} else {
				lex.push(newToken(str, IDENTIFIER))
			}
		} else if isLeftParenthese(lex.at()) {
			lex.push(newToken("(", LPAREN))
			lex.advanceN(1)
		} else if isRightParenthese(lex.at()) {
			lex.push(newToken(")", RPAREN))
			lex.advanceN(1)
		} else if isLeftBrace(lex.at()) {
			lex.push(newToken("{", LBRACE))
			lex.advanceN(1)
		} else if isRightBrace(lex.at()) {
			lex.push(newToken("}", RBRACE))
			lex.advanceN(1)
		} else if isLeftBracket(lex.at()) {
			lex.push(newToken("[", LBRACKET))
			lex.advanceN(1)
		} else if isRightBracket(lex.at()) {
			lex.push(newToken("]", RBRACKET))
			lex.advanceN(1)
		} else if isColon(lex.at()) {
			lex.push(newToken(":", COLON))
			lex.advanceN(1)
		} else if isOperator(lex.at()) {
			lex.push(newToken(string(lex.at()), OPERATOR))
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
