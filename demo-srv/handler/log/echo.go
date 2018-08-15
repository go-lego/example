package log

import (
	"strings"

	"github.com/go-lego/engine"
	"github.com/go-lego/engine/event"
	"github.com/go-lego/engine/log"
)

// Echo handler
type Echo struct {
}

// ID echo
func (e *Echo) ID() string {
	return "echo"
}

// OnSent handle echo.sent event
func (e *Echo) OnSent(ctx *engine.Context, evt *event.Event) error {
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
