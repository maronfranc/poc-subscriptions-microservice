package main

import (
	"github.com/maronfranc/subscription-system-products/src/config"
	"github.com/maronfranc/subscription-system-products/src/database"
	"github.com/maronfranc/subscription-system-products/src/server"
	"github.com/maronfranc/subscription-system-products/src/server/routes"
)

func main() {
	r := routes.Routes()
	cfg := config.GetConfig()
	dbcfg := database.Config{
		Username:     cfg.Database.Username,
		Password:     cfg.Database.Password,
		DatabaseName: cfg.Database.DatabaseName,
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
	}
	database.InitialiseDatabase(dbcfg)
	server.Listen(r, cfg.Server.Port)
}
