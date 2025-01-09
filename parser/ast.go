package parser

type NodeType string

const (
	CLASS        NodeType = "Class"
	RELATIONSHIP NodeType = "Relationship"
	INTERFACE    NodeType = "Interface"
	FIELD        NodeType = "Field"
	METHOD       NodeType = "Method"
	ROOT         NodeType = "Root"
)

type ASTNode struct {
	Type             NodeType
	RelationshipType string
	Name             string
	Visibility       string
	Parameters       []ASTNode
	Children         []ASTNode
	Default          string
	ValueType        string
	ReturnType       string
	SourceClass      string
	TargetClass      string
	LeftLabel        string
	MiddleLabel      string
	RightLabel       string
}
