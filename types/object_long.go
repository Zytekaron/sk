package types

import "strconv"

type Long struct {
	*Object
	Value int64
}

func NewLong(value int64) *Long {
	return &Long{
		Object: NewObject(),
		Value: value,
	}
}

func (l *Long) Type() Type {
	return "null"
}

func (l *Long) String() string {
	return strconv.FormatInt(l.Value, 10)
}

func (l *Long) Format(int, int, []Value) string {
	return l.String()
}

func (l *Long) Interface() interface{} {
	return l.String()
}
