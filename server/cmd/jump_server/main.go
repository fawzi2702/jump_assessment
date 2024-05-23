package main

import (
	"log"

	"github.com/this_is_iz/jump_server/internal/environement"
	"github.com/this_is_iz/jump_server/internal/server"
	"github.com/this_is_iz/jump_server/pkg/models"
)

func main() {
	// Load Env
	if envErr := environement.LoadEnv(); envErr != nil {
		log.Fatal(envErr)
	}

	// DB setup
	dbErr := models.InitializeDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	// Close DB connection
	defer models.CloseDB()

	// Setup server
	srv, srvErr := server.SetupServer()
	if srvErr != nil {
		panic(srvErr)
	}

	// Start server
	if srvErr = server.StartServer(srv); srvErr != nil {
		server.ShutdownServer(srv)
		panic(srvErr)
	}
	// Shutdown server
	defer server.ShutdownServer(srv)
}
