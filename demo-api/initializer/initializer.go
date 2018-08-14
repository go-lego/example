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
	bind.DefaultRegistry.Add("go.lego.srv.demo-srv", []*bind.Element{
		&bind.Element{ID: "echo.send"},
		&bind.Element{ID: "echo.sent"},
		&bind.Element{ID: "echo.error"},
	})
	api.AddService(&echo.Echo{})
}
