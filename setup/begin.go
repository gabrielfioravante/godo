package setup

import (
	"godo/db"
	"godo/options"
	"godo/server"
)

func Begin(opts *options.Flags) {
	db.Create()

	if opts.Migrate {
		db.Migrate()
	}

	server.Setup(opts.Mode, opts.Port)
}
