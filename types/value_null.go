package types

var null = &Null{}

type Null struct {}

func (n *Null) Type() Type {
	return "null"
}

func (n *Null) String() string {
	return "null"
}

func (n *Null) Format(int, int, []Value) string {
	return "null"
}

func (n *Null) Interface() interface{} {
	return n.String()
}