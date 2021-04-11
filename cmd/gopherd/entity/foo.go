package entity

type Foo struct {
	id int64
}

func NewFoo(id int64) *Foo {
	return &Foo{
		id: id,
	}
}

func (foo *Foo) ID() int64 {
	return foo.id
}
