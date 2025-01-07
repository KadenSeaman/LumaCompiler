package parser

import (
	"fmt"

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
	default:
		return nil, fmt.Errorf("unexpected token: %s", token.Value)
	}
}

func (p *Parser) parseClass() (*ASTNode, error) {
	// CLASS IDENTIFIER LBRACE

	p.nextToken() // skip class token
	nameToken := p.currentToken()

	if nameToken.Kind != lexer.IDENTIFIER {
		return nil, fmt.Errorf("expected class name, got :%s", nameToken.Value)
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
			return nil, fmt.Errorf("expected '}' got %s", p.currentToken().Value)
		}
		p.nextToken()
	}

	return classNode, nil
}

func (p *Parser) parseMember() (*ASTNode, error) {
	visibility := ""

	// Handle optional visibility
	if p.currentTokenKind() == lexer.OPERATOR {
		visibility = p.currentToken().Value
		p.nextToken() // skip visilbity
	}

	// Must have an identifier
	if p.currentTokenKind() != lexer.IDENTIFIER {
		return nil, fmt.Errorf("expected Identifier in property, got %s", p.currentToken().Value)
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
			return nil, fmt.Errorf("expected Identifier after colon in property, got %s", p.currentToken().Value)
		}
		valueType = p.currentToken().Value
		p.nextToken() // skip identifier

		//check for array value
		if p.currentTokenKind() == lexer.LBRACKET {
			valueType += p.currentToken().Value
			p.nextToken() // skip [

			if p.currentTokenKind() != lexer.RBRACKET {
				return nil, fmt.Errorf("expected ] after [ in field type decleration, got %s", p.currentToken().Value)
			}

			valueType += p.currentToken().Value
			p.nextToken() // skip ]
		}
	}

	if p.currentTokenKind() == lexer.EQUALS {
		p.nextToken() // skip equals
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, fmt.Errorf("expected Identifier after equals in property, got %s", p.currentToken().Value)
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
			return nil, fmt.Errorf("expected parameter name or ), got %s", p.currentToken().Value)
		}
	}

	if p.currentTokenKind() != lexer.RPAREN {
		return nil, fmt.Errorf("expected ) after method parameters, got %s", p.currentToken().Value)
	}
	p.nextToken() // skip )

	if p.currentTokenKind() == lexer.COLON {
		//specified type
		p.nextToken() // skip colon
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, fmt.Errorf("expected Identifier after colon in property, got %s", p.currentToken().Value)
		}

		//optional array return type

		returnType = p.currentToken().Value
		p.nextToken() //skip identifier

		if p.currentTokenKind() == lexer.LBRACKET {
			returnType += p.currentToken().Value
			p.nextToken() // skip [

			if p.currentTokenKind() != lexer.RBRACKET {
				return nil, fmt.Errorf("expected ] after [ in field type decleration, got %s", p.currentToken().Value)
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
		return nil, fmt.Errorf("expected identifier in interface, got %s", nameToken.Value)
	}
	interfaceNode := &ASTNode{Type: INTERFACE, Name: nameToken.Value}
	p.nextToken() // skip identifier

	return interfaceNode, nil
}
