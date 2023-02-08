package test

import (
	"belajar-go-restful-api-latihan2/app"
	"belajar-go-restful-api-latihan2/controller"
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/middleware"
	"belajar-go-restful-api-latihan2/model/domain"
	"belajar-go-restful-api-latihan2/repository"
	"belajar-go-restful-api-latihan2/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api_latihan2_test?parseTime=true")
	helper.PanicIfError(err)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)
	//meamtikan konkesi  apabila tidak digunakan lagi
	DB.SetConnMaxIdleTime(10 * time.Minute)

	//refresh coneksi kembali ke minimal database
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	repositorySiswa := repository.NewSiswaRepositoryImpl()
	repositoryGuru := repository.NewGuruRepositoryImpl()
	serviceSiswa := service.NewServiceSiswaRepositoryImpl(db, validate, repositorySiswa)
	serviceGuru := service.NewServiceGuruRepositoryImpl(db, repositoryGuru, validate)
	siswaController := controller.NewControllerSiswaImpl(serviceSiswa)
	guruController := controller.NewControllerGuruImpl(serviceGuru)
	router := app.NewRouter(siswaController, guruController)
	return middleware.NewAuthMiddleware(router)
}
func truncateSiswa(db *sql.DB) {
	db.Exec("TRUNCATE siswa")

}
func truncateGuru(db *sql.DB) {
	db.Exec("TRUNCATE guru")

}
func TestCreateSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name":"dono","nis":"12345","alamat":"medan"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/siswa", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "dono", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "12345", responseBody["data"].(map[string]interface{})["nis"])
}

func TestCreateSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name":"","nis":"12345","alamat":"medan"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/siswa", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestUpdateSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswa := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"dono","nis":"12345","alamat":"medan"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:2000/siswa/"+strconv.Itoa(siswa.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "dono", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "12345", responseBody["data"].(map[string]interface{})["nis"])
}

func TestUpdateSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswa := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"","nis":"12345","alamat":"medan"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:2000/siswa/"+strconv.Itoa(siswa.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestDeleteSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswa := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:2000/siswa/"+strconv.Itoa(siswa.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
}

func TestDeleteSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:2000/siswa/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}
func TestFindSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswa := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2000/siswa/"+strconv.Itoa(siswa.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "dono", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "12345", responseBody["data"].(map[string]interface{})["nis"])
}

func TestFindSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2000/siswa/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestFindAllSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	tx, err := db.Begin()
	siswaRepository := repository.NewSiswaRepositoryImpl()
	helper.PanicIfError(err)
	siswa1 := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "dono",
		Nis:    "12345",
		Alamat: "medan",
	})
	siswa2 := siswaRepository.Create(context.Background(), tx, domain.Siswa{
		Name:   "bima",
		Nis:    "123",
		Alamat: "berastagi",
	})
	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2000/siswa", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Println(responseBody)
	allSiswa := responseBody["data"].([]interface{})
	siswaResponse1 := allSiswa[0].(map[string]interface{})
	siswaResponse2 := allSiswa[1].(map[string]interface{})
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, siswa1.Id, int(siswaResponse1["id"].(float64)))
	assert.Equal(t, siswa1.Name, siswaResponse1["name"])
	assert.Equal(t, siswa1.Nis, siswaResponse1["nis"])

	assert.Equal(t, siswa2.Id, int(siswaResponse2["id"].(float64)))
	assert.Equal(t, siswa2.Name, siswaResponse2["name"])
	assert.Equal(t, siswa2.Nis, siswaResponse2["nis"])
}

func TestCreateGuruSuccess(t *testing.T) {
	db := setupTestDB()
	truncateGuru(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"id_guru":"GR01","name":"Bima","birth_day":"2023-02-06T00:00:00Z","nig":"12345","status":true}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/guru", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "Bima", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "GR01", responseBody["data"].(map[string]interface{})["id_guru"])
}
func TestCreateGuruFailed(t *testing.T) {
	db := setupTestDB()
	truncateGuru(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"id_guru":"","name":"Bima","birth_day":"2023-02-06T00:00:00Z","nig":"12345","status":true}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/guru", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("PAS-API-KEY", "MRGINZ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}
