package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"ctogram/internal/models"
	"ctogram/internal/router"
)

func (c *Cities) GetCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	city, err := c.Store.GetCityByID(id)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusNotFound, "404 page not found")
		return
	}

	output, err := json.Marshal(city)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError, "500 internal server error")
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
		ctx.WriteError(http.StatusInternalServerError, "500 internal server error")
		return
	}

	output, err := json.Marshal(cities)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError, "500 internal server error")
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
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	var city models.City
	if err := json.Unmarshal(b, &city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	if err := validateInput(city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Store.AddCity(city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError, "500 internal server error")
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	// ctx.ResponseWriter.Write([]byte("city added successfully"))
}

// DeleteCityHandler ..
func (c *Cities) DeleteCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	if err := c.Store.DeleteCityByID(id); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	// ctx.ResponseWriter.Write([]byte("city deleted successfully"))
}

// UpdateCityHandler ..
func (c *Cities) UpdateCityHandler(ctx *router.Context) {
	id, err := strconv.Atoi(ctx.Params["id"])
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	b, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	var city models.City
	if err := json.Unmarshal(b, &city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	if err := validateInput(city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Store.UpdateCity(id, city); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest, "400 bad request")
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	// ctx.ResponseWriter.Write([]byte("city updated successfully"))
}

func validateInput(city models.City) error {
	err := errors.New("400 invalid input data")
	switch false {
	case len(city.Name) > 0:
		return err
	case len(city.Code) > 0:
		return err
	case len(city.CountryCode) > 0:
		return err
	default:
		return nil
	}
}
