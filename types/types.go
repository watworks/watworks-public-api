package types

// Foo is the best type in the system!
type Foo struct {
	Name string
	Age  int16
}

// NewFoo returns a pointer to a Foo
func NewFoo() *Foo {
	return &Foo{}
}
