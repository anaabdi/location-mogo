package main

import (
	"log"
	"net/http"

	"flag"
	"fmt"
	"github.com/anaabdi/location-mogo/handler"
	"github.com/anaabdi/location-mogo/repository"
	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2"
)

func main() {
	mongoURL := flag.String("mongoURL", "localhost:27017", "connection url for mongodb")
	port := flag.String("port", "3222", "openend port for serving request")
	flag.Parse()

	session, err := mgo.Dial(*mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	areaRepo := repository.NewAreaRepo(session)
	areaHandler := handler.NewAreaHandler(areaRepo)
	router.Get("/api/v1/areas", areaHandler.GetByLocation())

	log.Printf("Starting to serve request to location-mogo in port %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), router))

}
