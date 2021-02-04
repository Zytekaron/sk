package types

// Value represents a value in sklang
type Value interface {
	// Type returns the type of this value
	Type() Type

	// String returns a string representation of this value
	String() string

	// Format formats the value into a pretty string
	Format(depth, offset int, visited []Value) string

	// Interface returns a interface{} for this type
	Interface() interface{}
}

type Iterable interface {
	Next() (Value, bool)
}
