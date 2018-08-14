package echo

import (
	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
)

// Echo service
type Echo struct {
}

// ID echo
func (e *Echo) ID() string {
	return "echo"
}

// OnSend handle echo.send event
func (e *Echo) OnSend(ng *engine.Engine, evt *engine.Event) error {
	log.Debug("Echo.OnSend is called")
	msg := evt.GetDataAsString("message")
	ng.AddResult("echo.message", msg)
	ne := engine.NewEvent("echo.sent", 0, evt)
	ne.SetData("message", msg)
	ng.RaiseEvent(ne)
	return nil
}

func (e *Echo) OnError(ng *engine.Engine, evt *engine.Event) error {
	return eerr.New(1111, "test")
}
