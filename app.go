package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/opensaucerer/gotemplate/config"

	"github.com/opensaucerer/gotemplate/middleware/v1"
	"github.com/opensaucerer/gotemplate/typing"
	"github.com/opensaucerer/gotemplate/version"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func createServer() (s *http.Server) {

	// create router
	r := mux.NewRouter()

	// wrap router into custom recover middleware
	rwr := middleware.Recover(r)

	// we should do more cross origin stuff here
	rc := middleware.CORS(rwr)

	// inject combined logger (apache & nginx style)
	logger := handlers.CombinedLoggingHandler(os.Stdout, rc)

	// register routes with versioning
	version.Version1Routes(r.StrictSlash(true))

	// load .env file
	env := config.MustGet("ENV_PATH", ".env")
	log.Printf("Loading %s file\n", env)
	if err := godotenv.Load(env); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading %s file\n", env)
		}
	}

	// verify env variables
	if err := config.VerifyEnvironment(typing.Env{}); err != nil {
		log.Fatalf("Error verifying environment variables: %s\n", err)
	}

	// append env variables to config.Env
	config.AppendEnvironment(config.Env)

	port := fmt.Sprintf(":%s", config.Env.Port)

	s = &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        logger,
	}

	go func() {
		log.Printf("Starting at http://127.0.0.1%s", port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening on port: %s\n", err)
		}
	}()

	return s
}

func main() {

	s := createServer()
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need to add it
	signal.Notify(config.ShutdownChan, syscall.SIGINT, syscall.SIGTERM)
	<-config.ShutdownChan
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shut down...")
	}
	log.Println("Server exited!")
}
