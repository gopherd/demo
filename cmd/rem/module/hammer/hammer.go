package hammer

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/rem/module"
	"github.com/gopherd/demo/cmd/rem/module/hammer/internal"
)

// Component is the interface that groups ther basic Component and HammerComponent methods
type Component interface {
	component.Component
	module.HammerComponent
}

// Component is the interface that groups ther basic Component and BarComponent methods
func NewComponent(service module.Service) Component {
	return internal.NewComponent(service)
}
