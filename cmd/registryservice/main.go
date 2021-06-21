package main

import (
	"context"
	"distributedServices/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	registry.SetupRegistryService()

	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServicePort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registery service started. Press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("shutting sown registry service")
}
