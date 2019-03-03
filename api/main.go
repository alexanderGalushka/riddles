package main

import (
	"fmt"
	consts "github.com/alexanderGalushka/riddles/api/constants"
	r "github.com/alexanderGalushka/riddles/api/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	router := r.InitRoutes()

	s := &http.Server{
		Addr:         consts.Port,
		Handler:      router,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}
	startupMsg := fmt.Sprintf("%s service is listening on port %s", consts.ServiceName, consts.Port)
	err := s.ListenAndServe()
	if err != nil {
		log.Print(startupMsg)
	} else {
		log.Fatal(fmt.Sprintf("failed to start %s web server", consts.ServiceName))
	}
}
