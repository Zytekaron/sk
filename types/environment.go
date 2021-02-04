package types

import (
	"fmt"
	"os"
)

type Environment struct {
	// data contains mappings from string to Value
	data map[string]Value

	// constants is a map for immutable identifiers
	constants map[string]bool

	// parent is the outer environment (for scope)
	parent *Environment
}

func NewEnv() *Environment {
	return &Environment{
		data:      make(map[string]Value),
		constants: make(map[string]bool),
		parent:    nil,
	}
}

func NewChildEnv(parent *Environment) *Environment {
	return &Environment{
		data:      make(map[string]Value),
		constants: make(map[string]bool),
		parent:    parent,
	}
}

// Get returns the Value an identifier stores
func (e *Environment) Get(name string) (Value, bool) {
	value, ok := e.data[name]
	if !ok && e.parent != nil {
		value, ok = e.parent.Get(name)
	}
	return value, ok
}

func (e *Environment) Declare(name string, value Value) Value {
	if _, ok := e.data[name]; ok {
		fmt.Println("Variable has already been declared:", name)
		os.Exit(1)
	}

	e.data[name] = value
	return value
}

func (e *Environment) DeclareConstant(name string, value Value) Value {
	if _, ok := e.data[name]; ok {
		fmt.Println("Variable has already been declared:", name)
		os.Exit(1)
	}

	e.data[name] = value
	e.constants[name] = true
	return value
}

// Set updates the value of a variable by name.
func (e *Environment) Set(name string, value Value) Value {
	if _, ok := e.data[name]; !ok {
		fmt.Printf("Cannot modify undeclared variable '%s'\n", name)
		os.Exit(1)
	}
	if constant, ok := e.constants[name]; ok && constant {
		fmt.Printf("Cannot modify constant variable '%s'\n", name)
		os.Exit(1)
	}

	e.data[name] = value
	return value
}
