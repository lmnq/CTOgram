package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ctogram/internal/handlers"
	"ctogram/internal/router"
	"ctogram/internal/store"
)

func main() {
	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	c := handlers.NewCities(s)

	r := router.NewRouter()
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	r.GET("/cities", c.GetCitiesHandler)
	r.POST("/cities", c.AddCityHandler)
	r.GET("/cities/:id", c.GetCityHandler)
	r.PUT("/cities/:id", c.UpdateCityHandler)
	r.DELETE("/cities/:id", c.DeleteCityHandler)

	fmt.Printf("server running on http://%v:%v\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), r))
}
