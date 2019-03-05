package main

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	r "github.com/alexanderGalushka/riddles/api/routers"
	"log"
	"net/http"
)

func main() {
	router := r.InitRoutes()

	s := &http.Server{
		Addr:         consts.Port,
		Handler:      router,
	}

	startupMsg := fmt.Sprintf("%s service is listening on port %s", consts.ServiceName, consts.Port)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start %s web server, error: %s", consts.ServiceName, err.Error()))
	} else {
		log.Print(startupMsg)
	}
}
