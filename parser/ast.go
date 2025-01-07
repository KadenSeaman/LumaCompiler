package parser

type NodeType int

const (
	CLASS NodeType = iota
	INTERFACE
	RELATIONSHIP
	FIELD
	METHOD
	ROOT
)

type ASTNode struct {
	Type       NodeType
	Name       string
	Visibility string
	Parameters []ASTNode
	Children   []ASTNode
}
