package main

import (
	"fmt"
	"github.com/ilborsch/go-execution-engine/internal/config"
	"github.com/ilborsch/go-execution-engine/internal/logger"
)

func main() {
	// init config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// init logger
	log := logger.SetupNew(cfg.Env)
	fmt.Println(log)
	// init docker conn

	// init request receiver & response sender (kafka or rest server)

	// init app
}
