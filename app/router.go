package app

import (
	"belajar-go-restful-api-latihan2/controller"
	"belajar-go-restful-api-latihan2/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(siswaController controller.ControllerSiswa, guruController controller.ControllerGuru) *httprouter.Router {
	router := httprouter.New()

	router.GET("/siswa", siswaController.FindAll)
	router.POST("/siswa", siswaController.Create)
	router.GET("/siswa/:siswaId", siswaController.FindById)
	router.PUT("/siswa/:siswaId", siswaController.Update)
	router.DELETE("/siswa/:siswaId", siswaController.Delete)

	router.GET("/guru", guruController.FindAll)
	router.POST("/guru", guruController.Create)
	router.GET("/guru/:guruId", guruController.FindById)
	router.PUT("/guru/:guruId", guruController.Update)
	router.DELETE("/guru/:guruId", guruController.Delete)
	router.PanicHandler = exception.ErrorHandler
	return router

}
