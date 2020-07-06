//Copyright (C) 2020 Larry Rau. all rights reserved
package parser

// Node is a parse tree element.
type Node interface {
	Type() NodeType
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// tree returns the containing *Tree.
	// It is unexported so all implementations of Node are in this package.
	tree() *Tree
}

// NodeType is the parse tree node ID
type NodeType int

const (
	NodeError      NodeType = iota //represent an error
	NodeIdentifier                 //an identifier
	NodeOption
	NodeInterface
	NodeAlias
	NodeFirewall
	NodeDefaults
	NodeNative
	Node
)
