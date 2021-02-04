package types

import "strconv"

type Int struct {
	*Object
	Value int32
}

func NewInt(value int32) *Int {
	return &Int{
		Object: NewObject(),
		Value: value,
	}
}

func (i *Int) Type() Type {
	return "int"
}

func (i *Int) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}

func (i *Int) Format(int, int, []Value) string {
	return i.String()
}

func (i *Int) Interface() interface{} {
	return i.String()
}
