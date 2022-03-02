package server

import (
	"context"
	"log"
	"oliversavio/api/middlewares"
	"oliversavio/api/routes"
	"oliversavio/api/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App      *fiber.App
	QuoteSvc service.Quotter
}

func (s *Server) setup() {

	if s.App == nil || s.QuoteSvc == nil {
		log.Fatalln("Server Incorrectly setup")
	}

	middlewares.FiberMiddlewares(s.App)
	routes.RegisterSwaggerRoute(s.App)
	routes.RegisterPublicRoutes(s.App, s.QuoteSvc)
	routes.RegisterPrometheusRoute(s.App)
}

func (s *Server) Start() <-chan os.Signal {
	s.setup()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.App.Listen(os.Getenv("SERVER_URL")); err != nil {
			log.Fatal(err)
		}
	}()
	return quit
}

func (s *Server) ShutdownGracefully() {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		// Release resources like Database connections
		cancel()
	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- s.App.Shutdown() }()

	select {
	case <-timeout.Done():
		log.Fatal("Server Shutdown Timed out before shutdown.")
	case err := <-shutdownChan:
		if err != nil {
			log.Fatal("Error while shutting down server", err)
		} else {
			log.Println("Server Shutdown Successful")
		}
	}
}
