package types

import "strconv"

var True = NewBool(true)
var False = NewBool(false)

type Bool struct {
	*Object
	Value bool
}

func NewBool(value bool) *Bool {
	return &Bool{Value: value}
}

func (i *Bool) Type() Type {
	return "null"
}

func (i *Bool) String() string {
	return strconv.FormatBool(i.Value)
}

func (i *Bool) Format(int, int, []Value) string {
	return i.String()
}

func (i *Bool) Interface() interface{} {
	return i.String()
}