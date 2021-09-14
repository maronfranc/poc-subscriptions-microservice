package main

import (
	"github.com/maronfranc/subscription-system-products/src/config"
	"github.com/maronfranc/subscription-system-products/src/mongodb"
	"github.com/maronfranc/subscription-system-products/src/server"
)

func main() {
	cfg := config.GetConfig()
	dbcfg := mongodb.Config{
		Username:     cfg.Database.Username,
		Password:     cfg.Database.Password,
		DatabaseName: cfg.Database.DatabaseName,
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
	}
	mongodb.InitialiseDatabase(dbcfg)
	server.Listen(cfg.Server.Port)
}
