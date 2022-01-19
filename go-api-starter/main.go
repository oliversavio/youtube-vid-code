package main

import (
	"context"
	"log"
	"oliversavio/api/middlewares"
	"oliversavio/api/routes"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "oliversavio/api/docs"

	"github.com/gofiber/fiber/v2"
)

// @title API Examples
// @version 1.0
// @description This is an auto-generated API Docs.
// @contact.name API Support
// @contact.email your@mail.com
// @BasePath /api
func main() {
	app := InitFiberApplication()

	middlewares.FiberMiddlewares(app)

	routes.RegisterSwaggerRoute(app)
	routes.RegisterPublicRoutes(app)

	waitforShutdownInterrupt := startServer(app)
	<-waitforShutdownInterrupt

	log.Println("Shutting Down Server..")
	shutdownGracefully(app)
}

func startServer(a *fiber.App) chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
			log.Fatal(err)
		}
	}()
	return quit
}

func shutdownGracefully(app *fiber.App) {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		// Release resources like Database connections
		cancel()
	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- app.Shutdown() }()

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
