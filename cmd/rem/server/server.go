package server

import (
	"time"

	"github.com/gopherd/doge/component"
	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/rem/module"
	"github.com/gopherd/demo/cmd/rem/module/hammer"
	"github.com/gopherd/demo/cmd/rem/module/horn"
	"github.com/gopherd/demo/pkg/config"
)

type server struct {
	*service.BaseApplication
	config *config.RemConfig

	quit, wait chan struct{}

	// components list all components of gopherd
	components struct {
		manager *component.Manager
		hammer  hammer.Component
		horn    horn.Component
		// more...
	}
}

// New creates application gopherd
func New() service.Application {
	s := &server{
		BaseApplication: service.NewBaseApplication(),
		config:          config.NewRemConfig(),
		quit:            make(chan struct{}),
		wait:            make(chan struct{}),
	}
	s.BaseApplication.SetConfigurator(s.config)

	s.components.manager = component.NewManager()
	s.components.hammer = s.components.manager.Add(hammer.NewComponent(s)).(hammer.Component)
	s.components.horn = s.components.manager.Add(horn.NewComponent(s)).(horn.Component)

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

func (s *server) GetHammerComponent() module.HammerComponent { return s.components.hammer }
func (s *server) GetHornComponent() module.HornComponent     { return s.components.horn }
