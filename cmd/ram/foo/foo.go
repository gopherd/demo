package foo

import (
	"github.com/gopherd/doge/service/component"

	"github.com/gopherd/demo/cmd/ram/foo/internal"
	"github.com/gopherd/demo/cmd/ram/module"
)

// NewComponent creates FooComponent
func NewComponent(service module.Service) interface {
	component.Component
	module.FooComponent
} {
	return internal.NewComponent(service)
}
