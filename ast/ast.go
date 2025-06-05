package ast

import (
	"strings"

	"github.com/timocheu/kalayo/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var builder strings.Builder

	for _, s := range p.Statements {
		builder.WriteString(s.String())
	}

	return builder.String()
}

// Start of nodes types

// Var holds the var token itself, identifier, value
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}
func (vs *VarStatement) String() string {
	var builder strings.Builder

	builder.WriteString(vs.TokenLiteral() + " ")
	builder.WriteString(vs.Name.String())
	builder.WriteString("=")

	if vs.Value != nil {
		builder.WriteString(vs.Value.String())
	}

	builder.WriteString(";")

	return builder.String()
}

// IDETIFER NODES
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Return statements
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
func (r *ReturnStatement) String() string {
	var builder strings.Builder

	builder.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		builder.WriteString(r.ReturnValue.String())
	}

	builder.WriteString(";")

	return builder.String()
}

// Expression statements
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}
func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}
func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(pe.Operator)
	builder.WriteString(pe.Right.String())
	builder.WriteString(")")

	return builder.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(ie.Left.String())
	builder.WriteString(" " + ie.Operator + " ")
	builder.WriteString(ie.Right.String())
	builder.WriteString(")")

	return builder.String()

}
