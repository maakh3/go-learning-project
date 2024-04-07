package main

import (
	"fmt"
	"net/http"

	"goapi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("\n---------\nStarting GO API service...")

	fmt.Println(`
________  ________          ________  ________  ___     
|\   ____\|\   __  \        |\   __  \|\   __  \|\  \    
\ \  \___|\ \  \|\  \       \ \  \|\  \ \  \|\  \ \  \   
 \ \  \  __\ \  \\\  \       \ \   __  \ \   ____\ \  \  
  \ \  \|\  \ \  \\\  \       \ \  \ \  \ \  \___|\ \  \ 
   \ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\
    \|_______|\|_______|        \|__|\|__|\|__|     \|__|`)

	fmt.Println("")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
