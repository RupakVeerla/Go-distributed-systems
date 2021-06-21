package main

import (
	"context"
	"distributedServices/log"
	"distributedServices/registry"
	"distributedServices/service"
	"fmt"

	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandler)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service!")
}
