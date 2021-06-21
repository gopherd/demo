package internal

import (
	"github.com/gopherd/doge/service/component"

	"github.com/gopherd/demo/cmd/rem/module"
)

// hornComponent implements horn.Component
type hornComponent struct {
	*component.BaseComponent

	service module.Service
}

func NewComponent(service module.Service) *hornComponent {
	return &hornComponent{
		BaseComponent: component.NewBaseComponent("horn"),
		service:       service,
	}
}

// Length implments HornComponent Length method
func (com *hornComponent) Length() int {
	return 222
}
