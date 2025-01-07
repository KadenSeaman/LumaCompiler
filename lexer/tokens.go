package lexer

import "fmt"

type Token struct {
	Value string
	Kind  TokenKind
}

// constructor
func newToken(value string, kind TokenKind) Token {
	return Token{Value: value, Kind: kind}
}

func tokenKindString(token Token) string {
	switch token.Kind {
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case CLASS:
		return "CLASS"
	case INTERFACE:
		return "INTERFACE"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LBRACKET:
		return "LBRACKET"
	case RBRACKET:
		return "RBRACKET"
	case OPERATOR:
		return "OPERATOR"
	case EQUALS:
		return "EQUALS"
	case SINGLE_LINE_COMMENT:
		return "COMMENT"
	case MULTI_LINE_COMMENT:
		return "MULTI_COMMENT"
	case ESCAPE:
		return "ESCAPE"
	case WHITESPACE:
		return "WHITESPACE"
	case NOT_FOUND:
		return "NOT_FOUND"
	case COLON:
		return "COLON"
	case COMMA:
		return "COMMA"
	default:
		return "UNKNOWN"
	}
}

var reservedLookup map[string]TokenKind = map[string]TokenKind{
	"class":     CLASS,
	"interface": INTERFACE,
}

var operatorLookup map[string]TokenKind = map[string]TokenKind{
	"<": OPERATOR,
	">": OPERATOR,
	"|": OPERATOR,
	"o": OPERATOR,
	"-": OPERATOR,
	".": OPERATOR,
	"*": OPERATOR,
	"x": OPERATOR,
	"#": OPERATOR,
	"~": OPERATOR,
	"/": OPERATOR,
	"+": OPERATOR,
}

func Debug(token Token) {
	fmt.Printf("%s('%s')\n", tokenKindString(token), token.Value)
}

type TokenKind int

const (
	// Special tokens
	EOF TokenKind = iota

	// Identifiers and literals
	IDENTIFIER

	// Keywords
	CLASS
	INTERFACE

	// Grouping
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
	COLON
	COMMA

	// Symbols & Operators
	OPERATOR
	EQUALS

	// Comments
	SINGLE_LINE_COMMENT
	MULTI_LINE_COMMENT
	ESCAPE

	// WHITESPACE
	WHITESPACE
	NOT_FOUND
)
