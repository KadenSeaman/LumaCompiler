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
		return "SINGLE_LINE_COMMENT"
	case OPEN_MULTI_LINE_COMMENT:
		return "OPEN_MULTI_LINE_COMMENT"
	case CLOSE_MULTI_LINE_COMMENT:
		return "CLOSE_MULTI_LINE_COMMENT"
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

func isReserved(str string) bool {
	switch str {
	case "class":
		return true
	case "void":
		return true
	case "null":
		return true
	case "interface":
		return true
	case "object":
		return true
	default:
		return false
	}
}

func getReservedTypeFromStr(str string) TokenKind {
	switch str {
	case "class":
		return CLASS
	case "void":
		return VOID
	case "null":
		return NULL
	case "interface":
		return INTERFACE
	case "object":
		return OBJECT
	default:
		return NOT_FOUND
	}
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
	OPEN_MULTI_LINE_COMMENT
	CLOSE_MULTI_LINE_COMMENT
	ESCAPE

	// WHITESPACE
	WHITESPACE
	NOT_FOUND
)
