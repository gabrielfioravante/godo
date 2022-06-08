package main

import (
	"flag"
	"godo/options"
	"godo/setup"
)

var flags options.Flags

func init() {
	flag.BoolVar(&flags.Migrate, "migrate", false, "should run migrations")
	flag.StringVar(&flags.Mode, "mode", "debug", "available modes: debug release test")
	flag.IntVar(&flags.Port, "port", 8080, "server port")
}

func main() {
	flag.Parse()

    setup.Begin(&flags)
}
