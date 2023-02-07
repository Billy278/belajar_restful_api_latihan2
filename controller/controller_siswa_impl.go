package controller

import (
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/web"
	"belajar-go-restful-api-latihan2/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ControllerSiswaImpl struct {
	service_siswa service.ServiceSiswaRepository
}

func NewControllerSiswaImpl(service_siswa service.ServiceSiswaRepository) ControllerSiswa {
	return &ControllerSiswaImpl{
		service_siswa: service_siswa,
	}
}

func (controller_siswa *ControllerSiswaImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	requestCreateSiswa := web.RequestCreateSiswa{}
	err := decoder.Decode(&requestCreateSiswa)
	helper.PanicIfError(err)
	responseSiswa := controller_siswa.service_siswa.Create(request.Context(), requestCreateSiswa)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseSiswa,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller_siswa *ControllerSiswaImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	requestUpdateSiswa := web.RequestUpdateSiswa{}
	helper.ReadFromRequest(request, &requestUpdateSiswa)

	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)
	requestUpdateSiswa.Id = id
	responseSiswa := controller_siswa.service_siswa.Update(request.Context(), requestUpdateSiswa)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseSiswa,
	}
	helper.WriteFromResponse(writer, &webResponse)

}

func (controller_siswa *ControllerSiswaImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)
	controller_siswa.service_siswa.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteFromResponse(writer, &webResponse)

}

func (controller_siswa *ControllerSiswaImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)
	responseSiswa := controller_siswa.service_siswa.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseSiswa,
	}
	helper.WriteFromResponse(writer, &webResponse)

}

func (controller_siswa *ControllerSiswaImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responseSiswa := controller_siswa.service_siswa.FindAll(request.Context())
	//fmt.Fprint(writer, responseSiswa)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseSiswa,
	}
	helper.WriteFromResponse(writer, &webResponse)
}
