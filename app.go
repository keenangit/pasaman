package application

import (
	"echo-demo-project/config"
	"log"

	"github.com/keenangit/pasaman/routes"
	"github.com/keenangit/pasaman/server"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
