package main

import (
	"belajar-go-restful-api-latihan2/app"
	"belajar-go-restful-api-latihan2/controller"
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/middleware"
	"belajar-go-restful-api-latihan2/repository"
	"belajar-go-restful-api-latihan2/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	repositorySiswa := repository.NewSiswaRepositoryImpl()
	repositoryGuru := repository.NewGuruRepositoryImpl()
	serviceSiswa := service.NewServiceSiswaRepositoryImpl(db, validate, repositorySiswa)
	serviceGuru := service.NewServiceGuruRepositoryImpl(db, repositoryGuru, validate)
	siswaController := controller.NewControllerSiswaImpl(serviceSiswa)
	guruController := controller.NewControllerGuruImpl(serviceGuru)

	// router := httprouter.New()

	// router.GET("/siswa", siswaController.FindAll)
	// router.POST("/siswa", siswaController.Create)
	// router.GET("/siswa/:siswaId", siswaController.FindById)
	// router.PUT("/siswa/:siswaId", siswaController.Update)
	// router.DELETE("/siswa/:siswaId", siswaController.Delete)

	// router.GET("/guru", guruController.FindAll)
	// router.POST("/guru", guruController.Create)
	// router.GET("/guru/:guruId", guruController.FindById)
	// router.PUT("/guru/:guruId", guruController.Update)
	// router.DELETE("/guru/:guruId", guruController.Delete)
	// router.PanicHandler = exception.ErrorHandler
	router := app.NewRouter(siswaController, guruController)

	//bisa juga begini
	// middlewate := middleware.AuthMiddleware{
	// 	Handler: router,
	// }
	// server := http.Server{
	// 	Addr:    "localhost:2000",
	// 	Handler: &middlewate,
	// }
	server := http.Server{
		Addr:    "localhost:2000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
