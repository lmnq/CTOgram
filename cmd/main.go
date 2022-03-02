package main

import (
	"ctogram/internal/handlers"
	"ctogram/internal/router"
	"ctogram/internal/store"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	c := handlers.NewCities(s)

	r := router.NewRouter()
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	r.GET("/cities", c.GetCitiesHandler)
	r.POST("/cities", c.AddCityHandler)
	r.GET("/cities/:id", c.GetCityHandler)
	r.PUT("/cities/:id", c.UpdateCityHandler)
	r.DELETE("/cities/:id", c.DeleteCityHandler)

	fmt.Printf("server running on http://%v:%v", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), r))
}
