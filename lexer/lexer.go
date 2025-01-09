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

func (lex *lexer) currentChar() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) charAtAdvanceN(n int) byte {
	return lex.source[lex.pos+n]
}

func (lex *lexer) isEOF() bool {
	return lex.pos >= len(lex.source)
}

func (lex *lexer) isEOFAtAdvanceN(n int) bool {
	return lex.pos+n >= len(lex.source)
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		tokens: make([]Token, 0),
	}
}

func (lexer *lexer) processAlphaNumeric() {
	var builder strings.Builder

	builder.WriteString(string(lexer.currentChar()))

	for !lexer.isEOFAtAdvanceN(1) && isAlphaNumeric(lexer.charAtAdvanceN(1)) {
		lexer.advanceN(1)
		builder.WriteString(string(lexer.currentChar()))
	}

	value := builder.String()

	if kind, exists := reservedLookup[value]; exists {
		lexer.push(newToken(value, kind))
	} else {
		lexer.push(newToken(value, IDENTIFIER))
	}
}

func (lexer *lexer) processSingleLineComment() {
	//comments
	lexer.advanceN(2)
	var builder strings.Builder

	for lexer.currentChar() != '\n' && !lexer.isEOF() {
		builder.WriteString(string(lexer.currentChar()))
		lexer.advanceN(1)
	}

	lexer.push(newToken(builder.String(), SINGLE_LINE_COMMENT))
}

func (lexer *lexer) processVisiblitySymbol() {
	kind := visibilityLookup[lexer.currentChar()]
	value := string(lexer.currentChar())

	lexer.push(newToken(value, kind))
}

func (lexer *lexer) processRelationship() {
	var builder strings.Builder

	builder.WriteString(string(lexer.currentChar()))

	for !lexer.isEOFAtAdvanceN(1) && isRelationshipSymbol(lexer.charAtAdvanceN(1)) {
		lexer.advanceN(1)
		builder.WriteString(string(lexer.currentChar()))
	}

	value := builder.String()

	if kind, exists := relationshipLookup[value]; exists {
		lexer.push(newToken(value, kind))
	} else if len(value) == 1 && isVisiblitySymbol([]byte(value)[0]) {
		lexer.processVisiblitySymbol()
	} else {
		lexer.handleUnknownString(value)
	}
}

func (lexer *lexer) processQuotation() {
	var builder strings.Builder

	for !lexer.isEOFAtAdvanceN(1) && !isQuotationMark(lexer.charAtAdvanceN(1)) {
		lexer.advanceN(1)
		builder.WriteString(string(lexer.currentChar()))
	}

	lexer.advanceN(1) // skip end quotes

	value := builder.String()

	lexer.push(newToken(value, QUOTATION))
}

func (lexer *lexer) handleUnknownCharacter() {
	fmt.Printf("Unknown character: %c\n", lexer.currentChar())
}

func (lexer *lexer) handleUnknownString(str string) {
	fmt.Printf("Unknown value: %s\n", str)
}

func Tokenize(source string) []Token {
	lexer := createLexer(source)

	// loop when there are still tokens remaining
	for !lexer.isEOF() {
		switch {
		case isWhitespace(lexer.currentChar()):
			lexer.advanceN(1)
			continue

		case isRelationshipSymbol(lexer.currentChar()):
			lexer.processRelationship()

		case isAlphaNumeric(lexer.currentChar()):
			lexer.processAlphaNumeric()

		case isLeftParenthese(lexer.currentChar()):
			lexer.push(newToken("(", LPAREN))

		case isRightParenthese(lexer.currentChar()):
			lexer.push(newToken(")", RPAREN))

		case isLeftBracket(lexer.currentChar()):
			lexer.push(newToken("[", LBRACKET))

		case isRightBracket(lexer.currentChar()):
			lexer.push(newToken("]", RBRACKET))

		case isLeftBrace(lexer.currentChar()):
			lexer.push(newToken("{", LBRACE))

		case isRightBrace(lexer.currentChar()):
			lexer.push(newToken("}", RBRACE))

		case isEquals(lexer.currentChar()):
			lexer.push(newToken("=", EQUALS))

		case isColon(lexer.currentChar()):
			lexer.push(newToken(":", COLON))

		case isComma(lexer.currentChar()):
			lexer.push(newToken(",", COMMA))

		case isVisiblitySymbol(lexer.currentChar()):
			lexer.processVisiblitySymbol()

		case isQuotationMark(lexer.currentChar()):
			lexer.processQuotation()

		case lexer.currentChar() == '/' && lexer.charAtAdvanceN(1) == '/':
			lexer.processSingleLineComment()

		default:
			lexer.handleUnknownCharacter()
		}

		lexer.advanceN(1)
	}

	lexer.tokens = append(lexer.tokens, newToken("EOF", EOF))
	return lexer.tokens
}
