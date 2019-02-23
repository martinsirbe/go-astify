package json

// Node represents a construct in an AST
type Node interface {
	Get() interface{}
}

// String represents a string node
type String struct {
	Value string
}

// Get implements Node interface
func (n String) Get() interface{} {
	return n
}

// Integer represents an integer node
type Integer struct {
	Value int
}

// Get implements Node interface
func (n Integer) Get() interface{} {
	return n
}

// Boolean represents a boolean node
type Boolean struct {
	Value bool
}

// Get implements Node interface
func (n Boolean) Get() interface{} {
	return n
}

// Array represents an array node
type Array struct {
	Elements []Node
}

// Get implements Node interface
func (n Array) Get() interface{} {
	return n
}

// Object represents a JSON object node
type Object struct {
	Properties map[string]Node
}

// Get implements Node interface
func (n Object) Get() interface{} {
	return n
}
