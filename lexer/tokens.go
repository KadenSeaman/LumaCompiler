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
	case VOID:
		return "VOID"
	case NULL:
		return "NULL"
	case INTERFACE:
		return "INTERFACE"
	case OBJECT:
		return "OBJECT"
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
	case SYMBOL:
		return "SYMBOL"
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
	default:
		return "UNKNOWN"
	}
}

var reservedLookup map[string]TokenKind = map[string]TokenKind{
	"class":     CLASS,
	"void":      VOID,
	"null":      NULL,
	"interface": INTERFACE,
	"object":    OBJECT,
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
	VOID
	NULL
	INTERFACE
	OBJECT

	// Grouping
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
	COLON

	// Symbols & Operators
	SYMBOL
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
