package initializer

import (
	"github.com/go-lego/engine/bind"
	"github.com/go-lego/engine/srv"
	"github.com/go-lego/example/demo-srv/handler/echo"
	"github.com/go-lego/example/demo-srv/handler/log"
)

type Initializer struct {
}

func (z *Initializer) InitServices() {
	srv.AddServices("echo", srv.PriorityLow, false, &echo.Echo{})
	srv.AddServices("log", srv.PriorityLow, false, &log.Echo{})

	bind.SetRegistry(bind.NewRedisRegistry("127.0.0.1:6379", "ttt"))
}
