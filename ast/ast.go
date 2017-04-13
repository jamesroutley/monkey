// Package ast implements Monkey's Abstract Syntax Tree
package ast

import (
	"bytes"
	"github.com/jamesroutley/monkey/token"
)

// Node is an item in the AST.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a Monkey statement in the AST.
// Statements are language constructs which don't produce values. Let, if and
// return are all statements.
type Statement interface {
	Node
	// Dummy method used to make the Go compiler distinguish Statements from
	// Expressions.
	statementNode()
}

// Expression represents a Monkey expression in the AST.
// Expressions are language constructs which produce values. Number or string
// literals and mathematical equations examples of expressions.
type Expression interface {
	Node
	// Dummy method used to make the Go compiler distinguish Expressions
	// from Statements.
	expressionNode()
}

// Program is the root of the AST and represents the program itself.
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement represents a variable binding statement in the AST.
// Let statements are denoted by the keyword 'let'.
type LetStatement struct {
	Token token.Token
	Name  *Identifier // Name of the variable being assigned to
	Value Expression  // Value of the variable being assigned to
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal value of the token associated with the
// statement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement represents a return statement in the AST.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the literal value of the return token.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type ExpressionStatement struct {
	// First token of the expression
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	} else {
		return ""
	}
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// Identifier represents a variable name in the AST.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the
// identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
