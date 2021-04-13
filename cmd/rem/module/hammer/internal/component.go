package internal

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/rem/module"
)

// hammerComponent implements hammer.Component
type hammerComponent struct {
	*component.BaseComponent

	service module.Service
}

func NewComponent(service module.Service) *hammerComponent {
	return &hammerComponent{
		BaseComponent: component.NewBaseComponent("hammer"),
		service:       service,
	}
}

// Init overrides Component Init method
func (com *hammerComponent) Init() error {
	if err := com.BaseComponent.Init(); err != nil {
		return err
	}
	// do something here...
	return nil
}

// Start overrides Component Start method
func (com *hammerComponent) Start() {
	com.BaseComponent.Start()
	// do something here...
}

// Shutdown overrides Component Shutdown method
func (com *hammerComponent) Shutdown() {
	// do something here...
	com.BaseComponent.Shutdown()
}

// Weight implments HammerComponent Weight method
func (com *hammerComponent) Weight() int {
	return 120
}
