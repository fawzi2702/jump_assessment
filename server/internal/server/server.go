package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/internal/environement"
	"github.com/this_is_iz/jump_server/pkg/routers"
)

type Server = http.Server

func SetupServer() (*Server, error) {
	var err error

	// Setup Gin
	r := gin.Default()

	// Setup CORS
	r.Use(cors.Default())

	// Set Gin mode
	envMode, err := environement.Get("MODE")
	if err != nil {
		return nil, err
	}

	if envMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Setup groups
	routers.SetupInvoiceRouter(r)
	routers.SetupUserRouter(r)
	routers.SetupTransactionRouter(r)

	// Start server
	apiPort, err := environement.Get("API_PORT")
	if err != nil {
		return nil, fmt.Errorf("API_PORT not found in environment: %w", err)
	}

	srv := &Server{
		Addr:    ":" + apiPort,
		Handler: r,
	}

	return srv, nil
}

func StartServer(srv *Server) error {
	if srv == nil {
		return fmt.Errorf("server is not set")
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	} else {
		fmt.Printf("Server started on %s\n", srv.Addr)

		return nil
	}
}

func ShutdownServer(srv *Server) {
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Server shutdown with error: %s\n", err)
	} else {
		fmt.Println("Server shutdown")
	}
}

func RestartServer(srv *Server) error {
	if err := srv.Shutdown(context.Background()); err != nil {
		return err
	}

	return StartServer(srv)
}
