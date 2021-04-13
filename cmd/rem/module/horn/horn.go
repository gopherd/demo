package horn

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/rem/module"
	"github.com/gopherd/demo/cmd/rem/module/horn/internal"
)

// Component is the interface that groups ther basic Component and HornComponent methods
type Component interface {
	component.Component
	module.HornComponent
}

// NewComponent returns a Component
func NewComponent(service module.Service) Component {
	return internal.NewComponent(service)
}
