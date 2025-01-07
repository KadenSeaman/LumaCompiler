package parser

type NodeType int
type RelationType int

const (
	CLASS NodeType = iota
	INTERFACE
	RELATIONSHIP
	FIELD
	METHOD
	ROOT
)

const (
	INHERITANCE RelationType = iota
	COMPOSITION
	AGGREGATION
	ASSOCIATION
)

type ASTNode struct {
	Type         NodeType
	RelationType RelationType
	Name         string
	Visibility   string
	Parameters   []ASTNode
	Children     []ASTNode
	Default      string
	ValueType    string
	ReturnType   string
	SourceClass  string
	TargetClass  string
}
