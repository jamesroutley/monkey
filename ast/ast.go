// Package ast implements Monkey's Abstract Syntax Tree
package ast

import (
	"github.com/jamesroutley/monkey/token"
)

// Node is an item in the AST.
type Node interface {
	TokenLiteral() string
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
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal value of the token associated with the
// statement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier represents a variable name in the AST.
type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the
// identifier
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
