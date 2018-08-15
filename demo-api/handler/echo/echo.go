package echo

import (
	"github.com/go-lego/engine"
	"github.com/go-lego/engine/api"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/event"
)

// Echo handler
type Echo struct {
}

// Name handler name
func (e *Echo) Name() string {
	return "echo"
}

// Routes handler routes
func (e *Echo) Routes() []*api.Route {
	return []*api.Route{
		&api.Route{Method: api.MethodPost, Path: "/", Handler: e.index, Input: &formIndex{}},
		&api.Route{Method: api.MethodPost, Path: "/error", Handler: e.error},
		&api.Route{Method: api.MethodGet, Path: "/rollback/{id}", Handler: e.rollback, Input: &formRollback{}},
	}
}

type formIndex struct {
	Message string `form:"msg" validate:"required"`
}

/**
 * @api {post} /echo Echo
 * @apiName Echo Message
 * @apiGroup Echo
 * @apiVersion 1.0.0
 * @apiDescription Echo a input message
 *
 * @apiParam {String} msg The input message
 *
 * @apiSuccess {Integer} code The result code
 * @apiSuccess {String} message The result message
 * @apiSuccess {String} data The echo message
 * @apiSuccessExample {json} Success Example:
	{
		"code": 0,
		"message": "Success",
		"data": "Hello world"
	}
 *
 * @apiUse ErrSystemError
 * @apiUse ErrInputInvalid
*/

func (e *Echo) index(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error) {
	ng := ctx.Engine()
	form := input.(*formIndex)
	evt := event.NewEvent("echo.send", 0, nil)
	evt.SetData("message", form.Message)
	ng.RaiseEvent(evt)
	if ng.HasError() {
		return nil, ng.Error(0)
	}
	return ng.Result("echo.message"), nil
}

/**
 * @api {post} /echo/error Error
 * @apiName Echo Error
 * @apiGroup Echo
 * @apiVersion 1.0.0
 * @apiDescription Test error
 *
 * @apiSuccess {Integer} code The result code
 * @apiSuccess {String} message The result message
 * @apiSuccessExample {json} Success Example:
	{
		"code": 0,
		"message": "Success"
	}
 *
 * @apiUse ErrSystemError
*/

func (e *Echo) error(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error) {
	ng := ctx.Engine()
	evt := event.NewEvent("echo.error", 0, nil)
	ng.RaiseEvent(evt)
	if ng.HasError() {
		return nil, ng.Error(0)
	}
	return "Error was expected, but success", nil
}

type formRollback struct {
	ID int `form:"id" validate:"required"`
}

/**
 * @api {post} /echo/rollback Rollback
 * @apiName Echo Rollback
 * @apiGroup Echo
 * @apiVersion 1.0.0
 * @apiDescription Test rollback
 *
 * @apiSuccess {Integer} code The result code
 * @apiSuccess {String} message The result message
 * @apiSuccessExample {json} Success Example:
	{
		"code": 0,
		"message": "Success"
	}
 *
 * @apiUse ErrSystemError
*/

func (e *Echo) rollback(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error) {
	form := input.(*formRollback)
	ng := ctx.Engine()
	evt := event.NewEvent("echo.rollback", 0, nil)
	evt.SetData("id", form.ID)
	ng.RaiseEvent(evt)
	if ng.HasError() {
		return nil, ng.Error(0)
	}
	return "Error was expected, but success", nil
}
