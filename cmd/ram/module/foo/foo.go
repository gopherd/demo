package foo

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/ram/module"
	"github.com/gopherd/demo/cmd/ram/module/foo/internal"
)

// Component is the interface that groups ther basic Component and FooComponent methods
type Component interface {
	component.Component
	module.FooComponent
}

// Component is the interface that groups ther basic Component and BarComponent methods
func NewComponent(service module.Service) Component {
	return internal.NewComponent(service)
}
