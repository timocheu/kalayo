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

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var builder strings.Builder

	builder.WriteString("if")
	builder.WriteString(ie.Condition.String())
	builder.WriteString(" ")
	builder.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		builder.WriteString("else ")
		builder.WriteString(ie.Alternative.String())
	}

	return builder.String()
}

type BlockStatement struct {
	Token      token.Token // '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var builder strings.Builder

	for _, s := range bs.Statements {
		builder.WriteString(s.String())
	}

	return builder.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var builder strings.Builder

	params := []string{}
	for _, s := range fl.Parameters {
		params = append(params, s.Value)
	}

	builder.WriteString(fl.TokenLiteral())
	builder.WriteString("(")
	builder.WriteString(strings.Join(params, ", "))
	builder.WriteString(")")
	builder.WriteString(fl.Body.String())

	return builder.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression // Could be Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var builder strings.Builder

	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	builder.WriteString(ce.Function.String())
	builder.WriteString("(")
	builder.WriteString(strings.Join(args, ", "))
	builder.WriteString(")")

	return builder.String()
}
