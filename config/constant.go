package config

import (
	"os"

	"github.com/opensaucerer/gotemplate/typing"
)

const (

	// EnvTagName is the tag name for environment variables struct
	envTagName = "env"

	// ShutdownTimeout is the time to wait for the server to shutdown gracefully
	ShutdownTimeout = 5 // seconds
)

var (
	// Env is the global environment variable
	Env = new(typing.Env) // global environment variable

	// ShutdownChan is the channel to listen for shutdown signals
	ShutdownChan = make(chan os.Signal, 1)
)
