package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"newsletter/config"
	netchi "newsletter/transport"
)

func main() {
	cfg, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	h := netchi.Initialize(cfg)

	fmt.Println(fmt.Sprintf("server is running on port: %d\n", cfg.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.Mux); err != nil {
		log.Fatal(err)
	}
}
