package initializer

import (
	"github.com/go-lego/engine/api"
	"github.com/go-lego/engine/bind"
	"github.com/go-lego/example/demo-api/handler/echo"
)

type Initializer struct {
}

func (z *Initializer) InitFilters() {
	// api.AddFilter(&filter.Auth{
	// 	Check: func(token string) (engine.Account, error) { return &MyAccount{}, nil },
	// 	Exclude: func() map[string]string {
	// 		return map[string]string{
	// 			"/v1/mobile/test1/echo": "*",
	// 		}
	// 	},
	// })
}

func (z *Initializer) InitServices() {
	api.AddService(&echo.Echo{})

	bind.SetRegistry(bind.NewRedisRegistry("127.0.0.1:6379", "ttt"))
}
