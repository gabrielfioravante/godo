package main

import (
	"flag"
	"godo/db"
	"godo/server"
)

func main() {
	db.Create()

	migrate := flag.Bool("migrate", false, "should run migrations")
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	if *migrate {
		db.Migrate()
	}

	server.Setup(*port)
}
