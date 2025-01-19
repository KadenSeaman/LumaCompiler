package parser

import (
	"errors"

	"github.com/kadenSeaman/lumaCompiler/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *Parser) nextToken() lexer.Token {
	if p.pos < len(p.tokens)-1 {
		p.pos++
	}

	return p.currentToken()
}

func (p *Parser) Parse() (*ASTNode, error) {
	root := &ASTNode{Type: ROOT, Name: "root"} // root node

	for p.currentTokenKind() != lexer.EOF {
		node, err := p.parseEntity()

		if err != nil {
			return nil, err
		}

		root.Children = append(root.Children, *node)
	}

	return root, nil
}

func (p *Parser) parseEntity() (*ASTNode, error) {
	token := p.currentToken()

	switch token.Kind {
	case lexer.CLASS:
		return p.parseClass()
	case lexer.INTERFACE:
		return p.parseInterface()
	//case for relationships
	case lexer.IDENTIFIER:
		return p.parseRelationship()
	default:
		return nil, errors.New("unexpected token: " + lexer.TokenKindName(token.Kind) + ", Token value: " + token.Value)
	}
}

func (p *Parser) parseRelationship() (*ASTNode, error) {
	sourceToken := p.currentToken()

	if sourceToken.Kind != lexer.IDENTIFIER {
		return nil, errors.New("expected token type: " + lexer.TokenKindName(sourceToken.Kind) + ", token value: " + sourceToken.Value + "in relationship, expected: IDENTIFIER")
	}

	relationshipNode := &ASTNode{Type: RELATIONSHIP, SourceClass: sourceToken.Value}

	p.nextToken() // skip identifier

	//optional left label here
	if p.currentToken().Kind == lexer.QUOTATION {
		relationshipNode.LeftLabel = p.currentToken().Value
		p.nextToken() // skip label
	}

	relationshipToken := p.currentToken()

	if !lexer.IsRelationshipKind(relationshipToken.Kind) {
		return nil, errors.New("expected relationship token, got token type of: " +
			lexer.TokenKindName(relationshipToken.Kind) +
			" and value of: " + relationshipToken.Value)
	}

	relationshipNode.RelationshipType = lexer.TokenKindName(relationshipToken.Kind)

	p.nextToken() //skip relationship

	//optional right label here
	if p.currentToken().Kind == lexer.QUOTATION {
		relationshipNode.RightLabel = p.currentToken().Value
		p.nextToken() // skip label
	}

	targetToken := p.currentToken()

	if targetToken.Kind != lexer.IDENTIFIER {
		return nil, errors.New("expected source class name in relationship, got: " + targetToken.Value)
	}

	relationshipNode.TargetClass = targetToken.Value

	p.nextToken() // skip relationship

	//optional middle label herre
	if p.currentToken().Kind == lexer.COLON {
		p.nextToken() // skip colon

		if p.currentToken().Kind != lexer.QUOTATION {
			return nil, errors.New("expected Quotation token for middle label on relationship, got token type of: " +
				lexer.TokenKindName(relationshipToken.Kind) + " and value of: " + relationshipToken.Value)
		}
		relationshipNode.MiddleLabel = p.currentToken().Value
		p.nextToken() // skip label
	}

	//remember to update name on relationship node
	relationshipNode.Name = sourceToken.Value + " " + relationshipToken.Value + " " + targetToken.Value

	return relationshipNode, nil
}

func (p *Parser) parseClass() (*ASTNode, error) {
	// CLASS IDENTIFIER LBRACE

	p.nextToken() // skip class token
	nameToken := p.currentToken()

	if nameToken.Kind != lexer.IDENTIFIER {
		return nil, errors.New("expected class name, got: " + nameToken.Value)
	}

	classNode := &ASTNode{Type: CLASS, Name: nameToken.Value}

	p.nextToken() // skip identifier

	if p.currentTokenKind() == lexer.LBRACE {
		p.nextToken()

		for p.currentTokenKind() != lexer.EOF && p.currentTokenKind() != lexer.RBRACE {
			propertyNode, err := p.parseMember()
			if err != nil {
				return nil, err
			}
			classNode.Children = append(classNode.Children, *propertyNode)
		}

		if p.currentTokenKind() != lexer.RBRACE {
			return nil, errors.New("expected '}', got: " + p.currentToken().Value)
		}
		p.nextToken() // skip }
	}

	return classNode, nil
}

func (p *Parser) parseMember() (*ASTNode, error) {
	visibility := ""

	// Handle optional visibility
	if p.currentTokenKind() == lexer.DASH ||
		p.currentTokenKind() == lexer.POUND ||
		p.currentTokenKind() == lexer.PLUS ||
		p.currentTokenKind() == lexer.TILDE {
		visibility = p.currentToken().Value
		p.nextToken() // skip visilbity
	}

	// Must have an identifier
	if p.currentTokenKind() != lexer.IDENTIFIER {
		return nil, errors.New("expected Identifier in property, got: " + p.currentToken().Value)
	}

	name := p.currentToken().Value
	p.nextToken() // Skip Identifier

	if p.currentTokenKind() == lexer.LPAREN {
		return p.parseMethod(name, visibility)
	}

	return p.parseField(name, visibility)
}

func (p *Parser) parseField(name string, visibility string) (*ASTNode, error) {
	valueType := ""
	defaultValue := ""

	if p.currentTokenKind() == lexer.COLON {
		//specified type
		p.nextToken() // skip colon
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, errors.New("expected Identifier after colon in property, got: " + p.currentToken().Value)
		}
		valueType = p.currentToken().Value
		p.nextToken() // skip identifier

		//check for array value
		if p.currentTokenKind() == lexer.LBRACKET {
			valueType += p.currentToken().Value
			p.nextToken() // skip [

			if p.currentTokenKind() != lexer.RBRACKET {
				return nil, errors.New("expected ']' after '[' in field type declaration, got: " + p.currentToken().Value)
			}

			valueType += p.currentToken().Value
			p.nextToken() // skip ]
		}
	}

	if p.currentTokenKind() == lexer.EQUALS {
		p.nextToken() // skip equals
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, errors.New("expected Identifier after equals in property, got: " + p.currentToken().Value)
		}
		defaultValue = p.currentToken().Value
		p.nextToken()
	}

	propertyNode := &ASTNode{Type: FIELD, Name: name, Visibility: visibility, ValueType: valueType, Default: defaultValue}

	return propertyNode, nil
}

func (p *Parser) parseMethod(name string, visibility string) (*ASTNode, error) {
	returnType := ""
	methodNode := &ASTNode{
		Type:       METHOD,
		Name:       name,
		Visibility: visibility,
	}

	p.nextToken() // skip (

	for p.currentTokenKind() != lexer.RPAREN && p.currentTokenKind() != lexer.EOF {
		if p.currentTokenKind() == lexer.IDENTIFIER {
			name := p.currentToken().Value
			p.nextToken()
			propertyNode, err := p.parseField(name, "")

			if err != nil {
				return nil, err
			}

			methodNode.Parameters = append(methodNode.Parameters, *propertyNode)
			if p.currentTokenKind() == lexer.COMMA {
				p.nextToken()
			}
		} else {
			return nil, errors.New("expected parameter name or ')', got: " + p.currentToken().Value)
		}
	}

	if p.currentTokenKind() != lexer.RPAREN {
		return nil, errors.New("expected ')' after method parameters, got: " + p.currentToken().Value)
	}
	p.nextToken() // skip )

	if p.currentTokenKind() == lexer.COLON {
		//specified type
		p.nextToken() // skip colon
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, errors.New("expected Identifier after colon in property, got: " + p.currentToken().Value)
		}

		//optional array return type

		returnType = p.currentToken().Value
		p.nextToken() //skip identifier

		if p.currentTokenKind() == lexer.LBRACKET {
			returnType += p.currentToken().Value
			p.nextToken() // skip [

			if p.currentTokenKind() != lexer.RBRACKET {
				return nil, errors.New("expected ']' after '[' in field type declaration, got: " + p.currentToken().Value)
			}

			returnType += p.currentToken().Value
			p.nextToken() // skip ]
		}

		methodNode.ReturnType = returnType
	}

	return methodNode, nil
}

func (p *Parser) parseInterface() (*ASTNode, error) {
	p.nextToken() // skip interface token

	nameToken := p.currentToken()

	if nameToken.Kind != lexer.IDENTIFIER {
		return nil, errors.New("expected identifier in interface, got: " + nameToken.Value)
	}
	interfaceNode := &ASTNode{Type: INTERFACE, Name: nameToken.Value}
	p.nextToken() // skip identifier

	return interfaceNode, nil
}
