package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type Null struct{}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type String struct {
	Value string
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []Object
}

// Type functions
func (i *Integer) Type() ObjectType      { return INTEGER_OBJ }
func (b *Boolean) Type() ObjectType      { return BOOLEAN_OBJ }
func (n *Null) Type() ObjectType         { return NULL_OBJ }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (e *Error) Type() ObjectType        { return ERROR_OBJ }
func (s *String) Type() ObjectType       { return STRING_OBJ }
func (b *Builtin) Type() ObjectType      { return BUILTIN_OBJ }
func (ao *Array) Type() ObjectType       { return ARRAY_OBJ }

// env
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// Inspect toString functions
func (i *Integer) Inspect() string      { return fmt.Sprintf("%d", i.Value) }
func (b *Boolean) Inspect() string      { return fmt.Sprintf("%t", b.Value) }
func (n *Null) Inspect() string         { return "null" }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (e *Error) Inspect() string        { return "ERROR: " + e.Message }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()

}
func (s *String) Inspect() string  { return s.Value }
func (b *Builtin) Inspect() string { return "builtin function" }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}

// Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	// env doesn't have the object, check outer env
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// enclosed Env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
