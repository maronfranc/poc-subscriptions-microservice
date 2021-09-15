package main

import (
	"github.com/maronfranc/subscription-system-products/src/config"
	"github.com/maronfranc/subscription-system-products/src/mongodb"
	"github.com/maronfranc/subscription-system-products/src/server"
)

func main() {
	config.Cfg = config.GetConfig()
	dbcfg := mongodb.Config{
		Username:     config.Cfg.Database.Username,
		Password:     config.Cfg.Database.Password,
		DatabaseName: config.Cfg.Database.DatabaseName,
		Host:         config.Cfg.Database.Host,
		Port:         config.Cfg.Database.Port,
	}
	mongodb.InitialiseDatabase(dbcfg)
	server.Listen(config.Cfg.Server.Port)
}
