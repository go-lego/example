package log

import (
	"strings"

	"github.com/go-lego/engine"
	"github.com/go-lego/engine/log"
)

type Echo struct {
}

func (e *Echo) ID() string {
	return "echo"
}

// OnSent handle echo.sent event
func (e *Echo) OnSent(ng *engine.Engine, evt *engine.Event) error {
	log.Debug("Echo.OnSent is called")
	log.Debug("Echoed message: %s", evt.GetDataAsString("message"))
	link := []string{}
	tt := evt
	for {
		if tt == nil {
			break
		}
		link = append(link, tt.Id)
		tt = tt.GetParent()
	}
	log.Debug("Event link:%s", strings.Join(link, "<="))
	return nil
}
