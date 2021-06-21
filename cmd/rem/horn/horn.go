package horn

import (
	"github.com/gopherd/doge/service/component"

	"github.com/gopherd/demo/cmd/rem/horn/internal"
	"github.com/gopherd/demo/cmd/rem/module"
)

// NewComponent creates HammerComponent
func NewComponent(service module.Service) interface {
	component.Component
	module.HornComponent
} {
	return internal.NewComponent(service)
}
