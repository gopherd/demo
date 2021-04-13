package bar

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/ram/module"
	"github.com/gopherd/demo/cmd/ram/module/bar/internal"
)

// Component is the interface that groups ther basic Component and BarComponent methods
type Component interface {
	component.Component
	module.BarComponent
}

// NewComponent returns a Component
func NewComponent(service module.Service) Component {
	return internal.NewComponent(service)
}
