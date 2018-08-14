package main

import (
	web "github.com/micro/go-web"

	"github.com/go-lego/engine/api"
	"github.com/go-lego/engine/log"
	"github.com/go-lego/example/demo-api/initializer"
)

// type testService struct{}

// func (ts *testService) Name() string {
// 	return "test1"
// }

// func (ts *testService) Routes() []*api.Route {
// 	return []*api.Route{
// 		&api.Route{Method: api.MethodGet, Path: "/echo", Handler: ts.echo, Input: &formEcho{}},
// 		&api.Route{Method: api.MethodGet, Path: "/info/{id}", Handler: ts.info, Input: &formInfo{}},
// 	}
// }

// type formEcho struct {
// 	Message string `form:"msg"`
// }

// func (ts *testService) echo(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error) {
// 	fmt.Println("Trying to echo the request data ....")
// 	form := input.(*formEcho)
// 	return "Message " + form.Message, nil
// }

// type formInfo struct {
// 	ID   int    `form:"id" validate:"required"`
// 	Name string `form:"name" validate:"required,email"`
// 	Test string `form:"test"`
// }

// func (ts *testService) info(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error) {
// 	fmt.Println("Trying to get info ....")
// 	form := input.(*formInfo)
// 	log.Debug("Form:%s", form)
// 	return fmt.Sprintf("Get info for %d", form.ID), nil
// }

// type initializer struct {
// }

// type MyAccount struct{}

// func (ma *MyAccount) ID() int64 { return 10 }
// func (z *initializer) InitFilters() {
// 	api.AddFilter(&filter.Auth{
// 		Check: func(token string) (engine.Account, error) { return &MyAccount{}, nil },
// 		Exclude: func() map[string]string {
// 			return map[string]string{
// 				"/v1/mobile/test1/echo": "*",
// 			}
// 		},
// 	})
// }

// func (z *initializer) InitServices() {
// 	s := time.Now()
// 	defer log.Profiling(s, "testService.InitServices")
// 	// time.Sleep(500 * time.Millisecond)
// 	api.AddService(&testService{})
// }

func main() {
	// New Service
	service := web.NewService(
		web.Name("go.lego.api.v1.mobile"),
		web.Version("latest"),
	)

	// Initialise service
	service.Init()

	if err := api.Init(service, new(initializer.Initializer)); err != nil {
		log.Fatal("%s", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("%s", err)
	}
}
