package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"newsletter/config"
	netchi "newsletter/transport"
)

func main() {
	configPath := "config.json"
	if !strings.Contains(os.Getenv("BASE_URL"), "localhost") {
		configPath = "../" + configPath
	}
	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	h := netchi.Initialize(cfg)

	fmt.Println(fmt.Sprintf("server is running on port: %d\n", cfg.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.Mux); err != nil {
		log.Fatal(err)
	}
}
