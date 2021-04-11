package internal

import (
	"time"

	"github.com/gopherd/doge/component"

	"github.com/gopherd/demo/cmd/gopherd/entity"
	"github.com/gopherd/demo/cmd/gopherd/module"
)

// fooComponent implements foo.Component
type fooComponent struct {
	*component.BaseComponent

	service module.Service
	foos    map[int64]*entity.Foo
}

func NewComponent(service module.Service) *fooComponent {
	return &fooComponent{
		BaseComponent: component.NewBaseComponent("foo"),
		service:       service,
		foos:          make(map[int64]*entity.Foo),
	}
}

// Init overrides Component Init method
func (com *fooComponent) Init() error {
	if err := com.BaseComponent.Init(); err != nil {
		return err
	}
	// do something here...
	return nil
}

// Start overrides Component Start method
func (com *fooComponent) Start() {
	com.BaseComponent.Start()
	// do something here...
}

// Shutdown overrides Component Shutdown method
func (com *fooComponent) Shutdown() {
	// do something here...
	com.BaseComponent.Shutdown()
}

// Update overrides Component Update method
func (com *fooComponent) Update(now time.Time, dt time.Duration) {
	for _, foo := range com.foos {
		com.onFooUpdate(now, dt, foo)
	}
}

// Get implments FooComponent Get method
func (com *fooComponent) Get(id int64) (*entity.Foo, bool) {
	foo, ok := com.foos[id]
	return foo, ok
}

// Remove implments FooComponent Remove method
func (com *fooComponent) Remove(id int64) *entity.Foo {
	if foo, ok := com.foos[id]; ok {
		delete(com.foos, id)
		return foo
	}
	return nil
}
