package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

var rootHandler = func(c *fiber.Ctx) error {
	return c.SendString("GO Fiber! Hello World")
}

var esInfoHandler = func(c *fiber.Ctx) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_IP"),
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error Creating Client %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Printf("Error getting info %s", err)
	}
	defer res.Body.Close()
	return c.SendString(res.String())
}

var pgConnectHandler = func(c *fiber.Ctx) error {
	log.Println("Connecting to DB")
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := pgx.Connect(timeoutCtx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}
	return c.SendString(greeting)
}
