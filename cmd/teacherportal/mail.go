package main

import (
	"context"
	"distributedServices/log"
	"distributedServices/registry"
	"distributedServices/service"
	"distributedServices/teacherportal"
	"fmt"
	stlog "log"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(), r, host, port, teacherportal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")

}
