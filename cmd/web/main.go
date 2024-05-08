package main

import (
	"fmt"
	"net/http"

	"github.com/wiscaksono/lapormas-go/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewChi(viperConfig)
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validator,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	addr := fmt.Sprintf(":%d", webPort)

	fmt.Printf("Starting server on %s\n", addr)
	http.ListenAndServe(addr, app)
}
