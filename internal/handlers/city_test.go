package handlers

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ctogram/internal/router"
	"ctogram/internal/store"
)

func TestCity(t *testing.T) {
	type test struct {
		url            string
		method         string
		expectedStatus int
		expectedData   []byte
		reqBody        []byte
	}

	type tester struct {
		name  string
		tests []test
	}

	tests := []tester{}

	tests = append(tests, tester{
		name: "testsAddCityHandler",
		tests: []test{
			{
				url:            "/cities",
				method:         http.MethodPost,
				expectedStatus: 201,
				reqBody:        []byte(`{"name":"Almaty","code":"ALM42","country_code":"911"}`),
			},
			{
				url:            "/cities",
				method:         http.MethodPost,
				expectedStatus: 201,
				reqBody:        []byte(`{"name":"Almaty2","code":"ALM42","country_code":"911"}`),
			},
			{
				url:            "/cities",
				method:         http.MethodPost,
				expectedStatus: 400,
				expectedData:   []byte(`{"Error":"400 invalid input data"}`),
				reqBody:        []byte(`{"invalidData":"false"}`),
			},
		},
	})

	tests = append(tests, tester{
		name: "testsGetCityHandler",
		tests: []test{
			{
				url:            "/cities/1",
				method:         http.MethodGet,
				expectedStatus: 200,
				expectedData:   []byte(`{"id":1,"name":"Almaty","code":"ALM42","country_code":"911"}`),
			},
			{
				url:            "/cities/404",
				method:         http.MethodGet,
				expectedStatus: 404,
				expectedData:   []byte(`{"Error":"404 page not found"}`),
			},
		},
	})

	tests = append(tests, tester{
		name: "testsGetCitiesHandler",
		tests: []test{
			{
				url:            "/cities",
				method:         http.MethodGet,
				expectedStatus: 200,
				expectedData:   []byte(`[{"id":1,"name":"Almaty","code":"ALM42","country_code":"911"},{"id":2,"name":"Almaty2","code":"ALM42","country_code":"911"}]`),
			},
		},
	})

	tests = append(tests, tester{
		name: "testsDeleteCityHandler",
		tests: []test{
			{
				url:            "/cities/2",
				method:         http.MethodDelete,
				expectedStatus: 200,
				expectedData:   []byte{},
			},
			{
				url:            "/cities/7",
				method:         http.MethodDelete,
				expectedStatus: 400,
				expectedData:   []byte(`{"Error":"400 bad request"}`),
			},
		},
	})

	tests = append(tests, tester{
		name: "testsUpdateCityHandler",
		tests: []test{
			{
				url:            "/cities/1",
				method:         http.MethodPut,
				expectedStatus: 200,
				reqBody:        []byte(`{"name":"Almaty1","code":"ALM11","country_code":"1111"}`),
			},
			{
				url:            "/cities/1",
				method:         http.MethodPut,
				expectedStatus: 400,
				expectedData:   []byte(`{"Error":"400 invalid input data"}`),
				reqBody:        []byte(`{"invalidName":"Almaty324","code":"ALM42142","country_code":"215911"}`),
			},
		},
	})

	tests = append(tests, tester{
		name: "testsGetCitiesAfterDeleteAndUpdate",
		tests: []test{
			{
				url:            "/cities",
				method:         http.MethodGet,
				expectedStatus: 200,
				expectedData:   []byte(`[{"id":1,"name":"Almaty1","code":"ALM11","country_code":"1111"}]`),
			},
		},
	})

	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	c := NewCities(s)
	t.Cleanup(func() {
		if err := CleanUpDB(c.Store.DB); err != nil {
			t.Error(err)
			return
		}
	})
	r := router.NewRouter()
	r.GET("/cities", c.GetCitiesHandler)
	r.POST("/cities", c.AddCityHandler)
	r.GET("/cities/:id", c.GetCityHandler)
	r.PUT("/cities/:id", c.UpdateCityHandler)
	r.DELETE("/cities/:id", c.DeleteCityHandler)

	for _, tl := range tests {
		for _, tc := range tl.tests {
			t.Run(tl.name, func(t *testing.T) {
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.reqBody))
				req.Header.Set("Content-Type", "application/json")
				r.ServeHTTP(rec, req)
				res := rec.Result()
				status := res.StatusCode
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Error(err)
				}
				defer res.Body.Close()
				if strings.TrimSpace(string(data)) != strings.TrimSpace(string(tc.expectedData)) {
					t.Error(tl.name)
					t.Errorf("\nexpected data: %v\ngot: %v\n", string(tc.expectedData), string(data))
				}
				if status != tc.expectedStatus {
					t.Error(tl.name)
					t.Errorf("\nexpected status: %v\ngot: %v\n", tc.expectedStatus, status)
				}
			})
		}
	}
}

// CleanUpDB ..
func CleanUpDB(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE cities`)
	if err != nil {
		return err
	}
	return store.CreateTable(db)
}
