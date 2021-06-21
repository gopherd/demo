package hammer

import (
	"github.com/gopherd/doge/service/component"

	"github.com/gopherd/demo/cmd/rem/hammer/internal"
	"github.com/gopherd/demo/cmd/rem/module"
)

// NewComponent creates HammerComponent
func NewComponent(service module.Service) interface {
	component.Component
	module.HammerComponent
} {
	return internal.NewComponent(service)
}
