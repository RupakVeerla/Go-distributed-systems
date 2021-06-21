package main

import (
	"context"
	"distributedServices/grades"
	"distributedServices/log"
	"distributedServices/registry"
	"distributedServices/service"
	"fmt"

	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegesterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		fmt.Printf("Logging service found at %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service!")
}
