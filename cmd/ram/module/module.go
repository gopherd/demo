package module

import (
	"github.com/gopherd/demo/cmd/ram/entity"
)

// Service represents a gopherd service
type Service interface {
	// GetFooComponent gets FooComponent
	GetFooComponent() FooComponent
	// GetBarComponent gets BarComponent
	GetBarComponent() BarComponent
}

// FooComponent represents foo component
type FooComponent interface {
	// Get gets entity Foo by id
	Get(id int64) (*entity.Foo, bool)
	// Remove removes entity Foo by id and return the removed entity
	Remove(id int64) *entity.Foo
}

// BarComponent represents bar component
type BarComponent interface {
	// Get gets entity Bar by id
	Get(id int64) (*entity.Bar, bool)
	// Remove removes entity Bar by id and return the removed entity
	Remove(id int64) *entity.Bar
}
