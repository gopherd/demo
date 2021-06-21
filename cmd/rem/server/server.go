package server

import (
	"time"

	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/rem/hammer"
	"github.com/gopherd/demo/cmd/rem/horn"
	"github.com/gopherd/demo/cmd/rem/module"
	"github.com/gopherd/demo/config"
)

type server struct {
	*service.BaseService
	config *config.RemConfig

	quit, wait chan struct{}

	// components list all components of rem
	components struct {
		hammer module.HammerComponent
		horn   module.HornComponent
		// more...
	}
}

// New creates rem service
func New() service.Service {
	cfg := config.NewRemConfig()
	s := &server{
		BaseService: service.NewBaseService(cfg),
		quit:        make(chan struct{}),
		wait:        make(chan struct{}),
	}

	s.components.hammer = s.BaseService.AddComponent(hammer.NewComponent(s)).(module.HammerComponent)
	s.components.horn = s.BaseService.AddComponent(horn.NewComponent(s)).(module.HornComponent)

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

func (s *server) GetHammerComponent() module.HammerComponent { return s.components.hammer }
func (s *server) GetHornComponent() module.HornComponent     { return s.components.horn }
