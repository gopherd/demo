package internal

import (
	"time"

	"github.com/gopherd/demo/cmd/gopherd/entity"
)

func (com *fooComponent) onFooUpdate(now time.Time, dt time.Duration, foo *entity.Foo) {
}
