package main

import (
	"github.com/go-lego/engine/log"
	"github.com/go-lego/engine/srv"
	"github.com/go-lego/example/demo-srv/initializer"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.lego.srv.demo-srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	if err := srv.Init(service, &initializer.Initializer{}); err != nil {
		log.Fatal("%s", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("%s", err)
	}
}
