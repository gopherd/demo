package bar

import (
	"github.com/gopherd/doge/service/component"

	"github.com/gopherd/demo/cmd/ram/bar/internal"
	"github.com/gopherd/demo/cmd/ram/module"
)

// NewComponent creates BarComponent
func NewComponent(service module.Service) interface {
	component.Component
	module.BarComponent
} {
	return internal.NewComponent(service)
}
