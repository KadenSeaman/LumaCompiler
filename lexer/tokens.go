package lexer

type Token struct {
	Value string
	Kind  TokenKind
}

// constructor
func newToken(value string, kind TokenKind) Token {
	return Token{Value: value, Kind: kind}
}

var reservedLookup map[string]TokenKind = map[string]TokenKind{
	"class":     CLASS,
	"interface": INTERFACE,
}

var visibilityLookup map[byte]TokenKind = map[byte]TokenKind{
	'-': DASH,
	'+': PLUS,
	'~': TILDE,
	'#': POUND,
}

var relationshipLookup map[string]TokenKind = map[string]TokenKind{
	"--":   ASSOCIATION,
	"<-->": BIDIR_ASSOCIATION,
	"-->":  R_ASSOCIATION,
	"<--":  L_ASSOCIATION,
	"..>":  R_DEPENDENCY,
	"<..":  L_DEPENDENCY,
	"--|>": R_INHERITANCE,
	"<|--": L_INHERITANCE,
	"..|>": R_IMPLEMENTATION,
	"<|..": L_IMPLEMENTATION,
	"--<>": R_AGGERGATION,
	"<>--": L_AGGREGATION,
	"--*":  R_COMPOSITION,
	"*--":  L_COMPOSITION,
}

func IsRelationshipKind(kind TokenKind) bool {
	switch kind {
	case ASSOCIATION,
		BIDIR_ASSOCIATION,
		R_ASSOCIATION,
		L_ASSOCIATION,
		R_DEPENDENCY,
		L_DEPENDENCY,
		R_INHERITANCE,
		L_INHERITANCE,
		R_IMPLEMENTATION,
		L_IMPLEMENTATION,
		R_AGGERGATION,
		L_AGGREGATION,
		R_COMPOSITION,
		L_COMPOSITION:
		return true
	default:
		return false
	}
}

func TokenKindName(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	// Identifiers and literals
	case IDENTIFIER:
		return "IDENTIFIER"
	// Keywords
	case CLASS:
		return "CLASS"
	case INTERFACE:
		return "INTERFACE"
	// Grouping
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
	case COLON:
		return "COLON"
	case COMMA:
		return "COMMA"
	case QUOTATION:
		return "QUOTATION"
	// Visibility
	case DASH:
		return "DASH"
	case PLUS:
		return "PLUS"
	case TILDE:
		return "TILDE"
	case POUND:
		return "POUND"
	// Relationships
	case ASSOCIATION:
		return "ASSOCIATION"
	case BIDIR_ASSOCIATION:
		return "BIDIR_ASSOCIATION"
	case R_ASSOCIATION:
		return "R_ASSOCIATION"
	case L_ASSOCIATION:
		return "L_ASSOCIATION"
	case R_DEPENDENCY:
		return "R_DEPENDENCY"
	case L_DEPENDENCY:
		return "L_DEPENDENCY"
	case R_INHERITANCE:
		return "R_INHERITANCE"
	case L_INHERITANCE:
		return "L_INHERITANCE"
	case R_IMPLEMENTATION:
		return "R_IMPLEMENTATION"
	case L_IMPLEMENTATION:
		return "L_IMPLEMENTATION"
	case R_AGGERGATION:
		return "R_AGGERGATION"
	case L_AGGREGATION:
		return "L_AGGREGATION"
	case R_COMPOSITION:
		return "R_COMPOSITION"
	case L_COMPOSITION:
		return "L_COMPOSITION"
	case EQUALS:
		return "EQUALS"
	// Comments
	case SINGLE_LINE_COMMENT:
		return "SINGLE_LINE_COMMENT"
	case MULTI_LINE_COMMENT:
		return "MULTI_LINE_COMMENT"
	case ESCAPE:
		return "ESCAPE"
	// Whitespace
	case WHITESPACE:
		return "WHITESPACE"
	case NOT_FOUND:
		return "NOT_FOUND"
	default:
		return "UNKNOWN"
	}
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
	QUOTATION

	// Visiblity
	DASH
	PLUS
	TILDE
	POUND

	// Relationships
	ASSOCIATION       // --
	BIDIR_ASSOCIATION // <-->
	R_ASSOCIATION     // -->
	L_ASSOCIATION     // <--
	R_DEPENDENCY      // ..>
	L_DEPENDENCY      // <..
	R_INHERITANCE     // --|>
	L_INHERITANCE     // <|--
	R_IMPLEMENTATION  // ..|>
	L_IMPLEMENTATION  // <|..
	R_AGGERGATION     // --o
	L_AGGREGATION     // o--
	R_COMPOSITION     // --*
	L_COMPOSITION     // *--

	EQUALS

	// Comments
	SINGLE_LINE_COMMENT
	MULTI_LINE_COMMENT
	ESCAPE

	// WHITESPACE
	WHITESPACE
	NOT_FOUND
)
