package service

import (
	"context"
	"distributedServices/registry"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, reg registry.Registration, host, port string, regesterHandlerFunc func()) (context.Context, error) {
	regesterHandlerFunc()
	ctx = startContext(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startContext(cxt context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(cxt)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
