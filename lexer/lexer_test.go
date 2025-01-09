package lexer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "class definition",
			input: "class MyClass { }",
			expected: []Token{
				{Value: "class", Kind: CLASS},
				{Value: "MyClass", Kind: IDENTIFIER},
				{Value: "{", Kind: LBRACE},
				{Value: "}", Kind: RBRACE},
				{Value: "EOF", Kind: EOF},
			},
		},
		{
			name:  "field with visibility",
			input: "- name: string",
			expected: []Token{
				{Value: "-", Kind: DASH},
				{Value: "name", Kind: IDENTIFIER},
				{Value: ":", Kind: COLON},
				{Value: "string", Kind: IDENTIFIER},
				{Value: "EOF", Kind: EOF},
			},
		},
		{
			name:  "relationship",
			input: "ClassA --> ClassB",
			expected: []Token{
				{Value: "ClassA", Kind: IDENTIFIER},
				{Value: "-->", Kind: R_ASSOCIATION},
				{Value: "ClassB", Kind: IDENTIFIER},
				{Value: "EOF", Kind: EOF},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := Tokenize(tt.input)
			if len(tokens) != len(tt.expected) {
				t.Errorf("Tokenize() got %d tokens, want %d tokens", len(tokens), len(tt.expected))
				return
			}

			for i, token := range tokens {
				if token.Kind != tt.expected[i].Kind || token.Value != tt.expected[i].Value {
					t.Errorf("Token[%d] = {%v, %v}, want {%v, %v}",
						i, token.Value, TokenKindName(token.Kind),
						tt.expected[i].Value, TokenKindName(tt.expected[i].Kind))
				}
			}
		})
	}
}
