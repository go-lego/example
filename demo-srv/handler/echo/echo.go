package echo

import (
	"errors"
	"strconv"

	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/event"
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
func (e *Echo) OnSend(ctx *engine.Context, evt *event.Event) error {
	log.Debug("Echo.OnSend is called")
	msg := evt.GetDataAsString("message")
	ctx.Engine().AddResult("echo.message", msg)
	ne := event.NewEvent("echo.sent", 0, evt)
	ne.SetData("message", msg)
	ctx.Engine().RaiseEvent(ne)
	return nil
}

// OnError handle echo.error event
func (e *Echo) OnError(ctx *engine.Context, evt *event.Event) error {
	return eerr.New(1111, "test")
}

// OnRollback handle echo.rollback event
func (e *Echo) OnRollback(ctx *engine.Context, evt *event.Event) error {
	log.Debug("Echo.OnRollback is called")
	id := evt.GetDataAsInt("id")
	ctx.Engine().AddResult("echo.rollback.id", strconv.Itoa(id))
	ne := event.NewEvent("echo.rolled", 0, evt)
	ctx.Engine().RaiseEvent(ne)
	return nil
}

// OnRolled handle echo.rolled event
func (e *Echo) OnRolled(ctx *engine.Context, evt *event.Event) error {
	return eerr.New(1111, "test")
}

// RollbackRollback rollback echo.rollback event
func (e *Echo) RollbackRollback(evt *event.Event, results map[string]string) error {
	log.Debug("Echo.RollbackRollback is called")
	log.Debug("Passed in results:%s", results)
	return errors.New("rollback failed")
}
