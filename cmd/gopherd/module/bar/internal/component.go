package internal

import (
	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/gopherd/entity"
	"github.com/gopherd/demo/cmd/gopherd/module"
)

// barComponent implements bar.Component
type barComponent struct {
	*component.BaseComponent

	service module.Service
	bars    map[int64]*entity.Bar
}

func NewComponent(service module.Service) *barComponent {
	return &barComponent{
		BaseComponent: component.NewBaseComponent("bar"),
		service:       service,
		bars:          make(map[int64]*entity.Bar),
	}
}

// Get implments BarComponent Get method
func (com *barComponent) Get(id int64) (*entity.Bar, bool) {
	bar, ok := com.bars[id]
	return bar, ok
}

// Remove implments BarComponent Remove method
func (com *barComponent) Remove(id int64) *entity.Bar {
	if bar, ok := com.bars[id]; ok {
		return bar
	}
	return nil
}
