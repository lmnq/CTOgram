package handlers

import (
	"ctogram/internal/models"
	"ctogram/internal/router"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (c *Cities) GetCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	city, err := c.Store.GetCityByID(id)
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "500 internal server error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(city)
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "500 internal server error", http.StatusInternalServerError)
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	ctx.ResponseWriter.Write(output)
}

// GetCitiesHandler ..
func (c *Cities) GetCitiesHandler(ctx *router.Context) {
	cities, err := c.Store.GetCities()
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "500 internal server error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(cities)
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "500 internal server error", http.StatusInternalServerError)
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	ctx.ResponseWriter.Write(output)
}

// AddCityHandler ..
func (c *Cities) AddCityHandler(ctx *router.Context) {
	b, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	var city models.City
	if err := json.Unmarshal(b, &city); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	if err := validateInput(city); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	if err := c.Store.AddCity(city); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "500 internal server error", http.StatusBadRequest)
		return
	}

	// ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	ctx.ResponseWriter.Write([]byte("city added successfully"))
}

// DeleteCityHandler ..
func (c *Cities) DeleteCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	if err := c.Store.DeleteCityByID(id); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	// ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	ctx.ResponseWriter.Write([]byte("city deleted successfully"))
}

// UpdateCityHandler ..
func (c *Cities) UpdateCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	var city models.City
	if err := json.Unmarshal(b, &city); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	if err := c.Store.UpdateCity(id, city); err != nil {
		log.Println(err)
		http.Error(ctx.ResponseWriter, "400 bad request", http.StatusBadRequest)
		return
	}

	// ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	ctx.ResponseWriter.Write([]byte("city updated successfully"))
}

func validateInput(city models.City) error {
	switch false {
	case len(city.Name) > 0:
		return errors.New("invalid input data")
	case len(city.Code) > 0:
		return errors.New("invalid input data")
	case len(city.CountryCode) > 0:
		return errors.New("invalid input data")
	default:
		return nil
	}
}
