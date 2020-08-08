package votearama

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Config contains all of the things you need to configure the server process
type Config struct {
	Port           int
	Host           string
	HealthEndpoint string
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("200 OK"))
}

// StartServer sets up the server and starts listening for traffic
func StartServer(config Config) {
	r := mux.NewRouter()

	r.HandleFunc(config.HealthEndpoint, handleHealth)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		loggedRouter,
	)

	if err != nil {
		log.Fatalf("Error while serving traffic: %s", err.Error())
	}
}
