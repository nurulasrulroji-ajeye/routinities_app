package test

import (
	"app/routinity/app"
	"app/routinity/controller"
	"app/routinity/helper"
	"app/routinity/middleware"
	"app/routinity/model/domain"
	"app/routinity/repository"
	"app/routinity/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/routinity_app")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	routinityRepository := repository.NewRoutintityRepo()
	routinityService := service.NewRoutinityService(routinityRepository, db, validate)
	routinityController := controller.NewRoutinityController(routinityService)
	router := app.NewRouter(routinityController)

	return middleware.NewAuthMiddleware(router)
}

func truncateRoutinity(db *sql.DB) {
	db.Exec("TRUNCATE db")
}

func TestCreateRoutinitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"activity" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/routinities", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["activity"])
}

func TestCreateRoutinityFailed(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"activity" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/routinities", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateRoutinitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)

	tx, _ := db.Begin()
	routinityRepository := repository.NewRoutintityRepo()
	routinity := routinityRepository.Save(context.Background(), tx, domain.Routinity{
		Activity: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/routinities/"+strconv.Itoa(routinity.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, routinity.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["activity"])
}

func TestUpdateRoutinityFailed(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)

	tx, _ := db.Begin()
	routinityRepository := repository.NewRoutintityRepo()
	routinity := routinityRepository.Save(context.Background(), tx, domain.Routinity{
		Activity: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/routinities/"+strconv.Itoa(routinity.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetRoutinitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)

	tx, _ := db.Begin()
	routinityRepository := repository.NewRoutintityRepo()
	routinity := routinityRepository.Save(context.Background(), tx, domain.Routinity{
		Activity: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/routinities/"+strconv.Itoa(routinity.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, routinity.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, routinity.Activity, responseBody["data"].(map[string]interface{})["activity"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateRoutinity(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/routinities/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}