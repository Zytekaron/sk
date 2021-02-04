package types

import "fmt"

type String struct {
	*Object
	Value string
}

func NewString(value string) *String {
	return &String{Value: value}
}

func (s *String) Type() Type {
	return "string"
}

func (s *String) String() string {
	return s.Value
}

func (s *String) Format(int, int, []Value) string {
	return fmt.Sprintf("\"%s\"", s.Value)
}

func (s *String) Interface() interface{} {
	return s.Value
}