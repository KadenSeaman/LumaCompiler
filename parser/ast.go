package parser

type NodeType int

const (
	CLASS NodeType = iota
	INTERFACE
	RELTATIONSHIP
	PROPERTY
)

type ASTNode struct {
	Type       NodeType
	Name       string
	Properties []ASTNode
	Children   []ASTNode
}
