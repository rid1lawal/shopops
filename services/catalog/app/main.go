package main

import (
	"log"

	"github.com/rid1lawal/shopops/services/catalog/internal/config"
	"github.com/rid1lawal/shopops/services/catalog/internal/server"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	log.Printf("catalog service listening on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
