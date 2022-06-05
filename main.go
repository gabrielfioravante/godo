package main

import (
	"flag"
	"godo/db"
	"godo/server"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Create()

	migrate := flag.Bool("migrate", false, "should run migrations")
	mode := flag.String("mode", "debug", "available modes: debug release test")
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	if *migrate {
		db.Migrate()
	}

	gin.SetMode(*mode)
	server.Setup(*port)
}
