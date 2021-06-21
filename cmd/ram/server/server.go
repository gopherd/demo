package server

import (
	"time"

	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/ram/bar"
	"github.com/gopherd/demo/cmd/ram/foo"
	"github.com/gopherd/demo/cmd/ram/module"
	"github.com/gopherd/demo/config"
)

type server struct {
	*service.BaseService
	config *config.RamConfig

	quit, wait chan struct{}

	// components list all components of ram
	components struct {
		foo module.FooComponent
		bar module.BarComponent
		// more...
	}
}

// New creates ram service
func New() service.Service {
	cfg := config.NewRamConfig()
	s := &server{
		BaseService: service.NewBaseService(cfg),
		quit:        make(chan struct{}),
		wait:        make(chan struct{}),
	}

	s.components.foo = s.AddComponent(foo.NewComponent(s)).(module.FooComponent)
	s.components.bar = s.AddComponent(bar.NewComponent(s)).(module.BarComponent)

	return s
}

// Init overrides BaseService Init method
func (s *server) Init() error {
	return s.BaseService.Init()
}

// Start overrides BaseService Start method
func (s *server) Start() error {
	s.BaseService.Start()
	go s.run()
	return nil
}

// Shutdown overrides BaseService Shutdown method
func (s *server) Shutdown() error {
	close(s.quit)
	<-s.wait
	s.BaseService.Shutdown()
	return nil
}

// run runs service's main loop
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
	s.BaseService.Update(now, dt)
}

// implements interface module.Service

func (s *server) GetFooComponent() module.FooComponent { return s.components.foo }
func (s *server) GetBarComponent() module.BarComponent { return s.components.bar }
