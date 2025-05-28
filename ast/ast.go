package ast

import "github.com/timocheu/kalayo/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	// takes an node interface
	Node
	statementNode()
}

type Expression interface {
	// takes an node interface
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	// return the first literal in the "Statements"
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Start of nodes types

// Var holds the var token itself, identifier, value
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *VarStatement) statementNode() {}
func (ls *VarStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
