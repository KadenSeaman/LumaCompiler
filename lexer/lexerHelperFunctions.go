package lexer

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

func isComma(c byte) bool {
	return c == ','
}

func isEquals(c byte) bool {
	return c == '='
}

func isVisiblitySymbol(c byte) bool {
	if _, exists := visibilityLookup[c]; exists {
		return true
	} else {
		return false
	}
}

func isRelationshipSymbol(c byte) bool {
	if c == '-' || c == '.' || c == '>' || c == '<' || c == '|' || c == '*' {
		return true
	} else {
		return false
	}
}

func isQuotationMark(c byte) bool {
	if c == '"' {
		return true
	} else {
		return false
	}
}
