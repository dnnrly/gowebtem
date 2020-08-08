package main

import (
	"flag"

	votearama "github.com/dnnrly/gowebtem"
)

func main() {
	host := flag.String("host", "", "listening host address")
	port := flag.Int("port", 8080, "listening port")
	healthEndpoint := flag.String("health", "/health", "health endpoint location")

	config := votearama.Config{
		Host:           *host,
		Port:           *port,
		HealthEndpoint: *healthEndpoint,
	}

	votearama.StartServer(config)
}
