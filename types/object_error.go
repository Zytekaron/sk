package types

import (
	"fmt"
	"os"
)

type Error struct {
	Name        string
	Description string
}

func NewError(name, description string) *Error {
	return &Error{
		Name:        name,
		Description: description,
	}
}

func (e *Error) Panic() {
	fmt.Println(e.String())
	os.Exit(1)
}

func (e *Error) Type() Type {
	return "error"
}

func (e *Error) String() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Description)
}

func (e *Error) Format(depth, offset int, visited []Value) string {
	obj := NewObject()
	obj.Set("name", NewString(e.Name))
	obj.Set("description", NewString(e.Description))
	return obj.Format(depth, offset, visited)
}

func (e *Error) Call(_ *Environment, _ string, _ ...Value) Value {
	return null // fixme fixme fixme
}

func (e *Error) Interface() interface{} {
	return e.String()
}
