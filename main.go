package main

import (
	"aquarium/routers"
	"net/http"
	"aquarium/tool/log"
)

func main() {
	if err := config.InitConfig(); err != nil{
		log.Fatalf("init config failed",err)
	}
	router := routers.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handle
	}	
	
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
