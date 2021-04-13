package server

import (
	"time"

	"github.com/gopherd/doge/component"
	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/ram/module"
	"github.com/gopherd/demo/cmd/ram/module/bar"
	"github.com/gopherd/demo/cmd/ram/module/foo"
	"github.com/gopherd/demo/pkg/config"
)

type server struct {
	*service.BaseApplication
	config *config.RamConfig

	quit, wait chan struct{}

	// components list all components of gopherd
	components struct {
		manager *component.Manager
		foo     foo.Component
		bar     bar.Component
		// more...
	}
}

// New creates application gopherd
func New() service.Application {
	s := &server{
		BaseApplication: service.NewBaseApplication(),
		config:          config.NewRamConfig(),
		quit:            make(chan struct{}),
		wait:            make(chan struct{}),
	}
	s.BaseApplication.SetConfigurator(s.config)

	s.components.manager = component.NewManager()
	s.components.foo = s.components.manager.Add(foo.NewComponent(s)).(foo.Component)
	s.components.bar = s.components.manager.Add(bar.NewComponent(s)).(bar.Component)

	return s
}

// Init overrides BaseApplication Init method
func (s *server) Init() error {
	return s.components.manager.Init()
}

// Start overrides BaseApplication Start method
func (s *server) Start() error {
	s.components.manager.Start()
	go s.run()
	return nil
}

// Shutdown overrides BaseApplication Shutdown method
func (s *server) Shutdown() error {
	close(s.quit)
	<-s.wait
	s.components.manager.Shutdown()
	return nil
}

// run runs application's main loop
func (s *server) run() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	lastUpdatedAt := time.Now()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			s.onUpdate(now, now.Sub(lastUpdatedAt))
			lastUpdatedAt = now
		case <-s.quit:
			close(s.wait)
			return
		}
	}
}

func (s *server) onUpdate(now time.Time, dt time.Duration) {
	s.components.manager.Update(now, dt)
}

// implements interface module.Service

func (s *server) GetFooComponent() module.FooComponent { return s.components.foo }
func (s *server) GetBarComponent() module.BarComponent { return s.components.bar }
