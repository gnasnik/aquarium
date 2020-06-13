package main

import (
	"aquarium/config"
	"aquarium/routers"
	"aquarium/utils/env"
	glog "aquarium/utils/log"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal("init config failed, %v", err)
	}

	logLevel := "info"
	if config.Configs.RunMode == "debug" {
		logLevel = "debug"
	}

	logger, _ := glog.NewLogger(env.LogPath, logLevel)
	glog.SetDefault(logger)

	router := routers.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen: %s\n", err)
	}
}
