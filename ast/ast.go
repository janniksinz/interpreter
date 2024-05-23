package ast

import (
	"bytes"
	"monkey/token"
)

// STRUCTS

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement s are like variable bindings // let
// they don't produce values
type Statement interface {
	Node
	statementNode()
}

// Expression produce values, Statement doesn't
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// Identifier for LetStatement
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// ExpressionStatement /*
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// IntegerLiteral Expression
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // the prefix token, e.g. ! -
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// Interface methods for
//	- ReturnStatement
//	- LetStatement
//	- Identifier
//	- ExpressionStatement
//	- IntegerLiteral	*/

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String implementations for
//   - Program
//   - LetStatement
//   - ReturnStatement
//   - ExpressionStatement
//   - IntegerLiteral
//   - Identifier*/
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (il *IntegerLiteral) String() string { return il.Token.Literal }

func (i *Identifier) String() string { return i.Value }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}