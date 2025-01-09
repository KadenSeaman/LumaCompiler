package parser

import (
	"testing"

	"github.com/kadenSeaman/lumaCompiler/lexer"
)

func TestParseClass(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ASTNode
	}{
		{
			name:  "empty class",
			input: "class MyClass { }",
			expected: &ASTNode{
				Type: CLASS,
				Name: "MyClass",
			},
		},
		{
			name: "class with field",
			input: `class MyClass {
				- name: string
			}`,
			expected: &ASTNode{
				Type: CLASS,
				Name: "MyClass",
				Children: []ASTNode{
					{
						Type:       FIELD,
						Name:       "name",
						Visibility: "-",
						ValueType:  "string",
					},
				},
			},
		},
		{
			name: "class with method",
			input: `class MyClass {
				+ getName(): string
			}`,
			expected: &ASTNode{
				Type: CLASS,
				Name: "MyClass",
				Children: []ASTNode{
					{
						Type:       METHOD,
						Name:       "getName",
						Visibility: "+",
						ReturnType: "string",
					},
				},
			},
		},
		{
			name: "class with method and paramaters",
			input: `class MyClass {
				+ getName(str : string): string
			}`,
			expected: &ASTNode{
				Type: CLASS,
				Name: "MyClass",
				Children: []ASTNode{
					{
						Type:       METHOD,
						Name:       "getName",
						Visibility: "+",
						ReturnType: "string",
						Parameters: []ASTNode{
							{
								Type:      FIELD,
								Name:      "str",
								ValueType: "string",
							},
						},
					},
				},
			},
		},
		{
			name: "class with method and paramaters with default values",
			input: `class MyClass {
				+ getName(str : string = example): string
			}`,
			expected: &ASTNode{
				Type: CLASS,
				Name: "MyClass",
				Children: []ASTNode{
					{
						Type:       METHOD,
						Name:       "getName",
						Visibility: "+",
						ReturnType: "string",
						Parameters: []ASTNode{
							{
								Type:      FIELD,
								Name:      "str",
								ValueType: "string",
								Default:   "example",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := lexer.Tokenize(tt.input)
			p := NewParser(tokens)
			ast, err := p.Parse()

			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}

			// Since Parse() returns a root node, we need to check its first child
			if len(ast.Children) == 0 {
				t.Errorf("Parse() returned empty AST")
				return
			}

			result := &ast.Children[0]
			compareNodes(t, result, tt.expected)
		})
	}
}

func TestParseRelationship(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ASTNode
	}{
		{
			name:  "simple association",
			input: "ClassA --> ClassB",
			expected: &ASTNode{
				Type:             RELATIONSHIP,
				RelationshipType: "R_ASSOCIATION",
				Name:             "ClassA --> ClassB",
				SourceClass:      "ClassA",
				TargetClass:      "ClassB",
			},
		},
		{
			name:  "labeled relationship",
			input: `ClassA "1" --> "many" ClassB: "contains"`,
			expected: &ASTNode{
				Type:             RELATIONSHIP,
				RelationshipType: "R_ASSOCIATION",
				Name:             "ClassA --> ClassB",
				SourceClass:      "ClassA",
				TargetClass:      "ClassB",
				LeftLabel:        "1",
				RightLabel:       "many",
				MiddleLabel:      "contains",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := lexer.Tokenize(tt.input)
			p := NewParser(tokens)
			ast, err := p.Parse()

			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}

			if len(ast.Children) == 0 {
				t.Errorf("Parse() returned empty AST")
				return
			}

			result := &ast.Children[0]
			compareNodes(t, result, tt.expected)
		})
	}
}

// Helper function to compare AST nodes
func compareNodes(t *testing.T, got, want *ASTNode) {
	if got.Type != want.Type {
		t.Errorf("Node.Type = %v, want %v", got.Type, want.Type)
	}
	if got.Name != want.Name {
		t.Errorf("Node.Name = %v, want %v", got.Name, want.Name)
	}
	if got.Visibility != want.Visibility {
		t.Errorf("Node.Visibility = %v, want %v", got.Visibility, want.Visibility)
	}
	if got.ValueType != want.ValueType {
		t.Errorf("Node.ValueType = %v, want %v", got.ValueType, want.ValueType)
	}
	if got.ReturnType != want.ReturnType {
		t.Errorf("Node.ReturnType = %v, want %v", got.ReturnType, want.ReturnType)
	}

	if len(got.Children) != len(want.Children) {
		t.Errorf("Node.Children length = %d, want %d", len(got.Children), len(want.Children))
		return
	}

	for i := range got.Children {
		compareNodes(t, &got.Children[i], &want.Children[i])
	}
}
