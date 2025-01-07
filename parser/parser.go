package parser

import (
	"fmt"

	"github.com/kadenSeaman/lumaCompiler/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func newParser(tokens []lexer.Token) *Parser {
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

func (p *Parser) parse() (*ASTNode, error) {
	root := &ASTNode{Type: PROPERTY, Name: "root"} // root node

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

	if token.Kind == lexer.CLASS {
		return p.parseClass()
	} else if token.Kind == lexer.INTERFACE {
		return p.parseInterface()
	}

	return nil, fmt.Errorf("unexpected token: %s", token.Value)
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
			propertyNode, err := p.parseProperty()
			if err != nil {
				return nil, err
			}
			classNode.Properties = append(classNode.Properties, *propertyNode)
		}

		if p.currentTokenKind() != lexer.RBRACE {
			return nil, fmt.Errorf("expected '}' got %s", p.currentToken().Value)
		}
		p.nextToken()
	}

	return classNode, nil
}

func (p *Parser) parseProperty() (*ASTNode, error) {
	if p.currentTokenKind() != lexer.OPERATOR {
		return nil, fmt.Errorf("expected visiblity in property, got %s", p.currentToken().Value)
	}
	visibilityToken := p.currentToken()
	p.nextToken() // skip visilbity

	if p.currentTokenKind() != lexer.IDENTIFIER {
		return nil, fmt.Errorf("expected Identifier in property, got %s", p.currentToken().Value)
	}
	identifierToken := p.currentToken()

	propertyValue := visibilityToken.Value + identifierToken.Value

	propertyNode := &ASTNode{Type: PROPERTY, Name: propertyValue}
	p.nextToken() // skip identifier
	return propertyNode, nil
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
