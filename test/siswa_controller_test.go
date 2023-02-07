package test

import (
	"belajar-go-restful-api-latihan2/app"
	"belajar-go-restful-api-latihan2/controller"
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/middleware"
	"belajar-go-restful-api-latihan2/repository"
	"belajar-go-restful-api-latihan2/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
func TestCreateSiswa(t *testing.T) {
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
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "dono", responseBody["data"].(map[string]interface{})["name"])

}
