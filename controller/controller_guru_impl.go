package controller

import (
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/web"
	"belajar-go-restful-api-latihan2/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerGuruImpl struct {
	Service_guru service.ServiceGuruRepository
}

func NewControllerGuruImpl(service_guru service.ServiceGuruRepository) ControllerGuru {
	return &ControllerGuruImpl{
		Service_guru: service_guru,
	}
}
func (control_guru *ControllerGuruImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestCreateguru := web.RequestCreateGuru{}
	helper.ReadFromRequest(request, &requestCreateguru)
	responseGuru := control_guru.Service_guru.Create(request.Context(), requestCreateguru)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseGuru,
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control_guru *ControllerGuruImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	guruId := params.ByName("guruId")
	requestUpdateguru := web.RequestUpdateGuru{}
	requestUpdateguru.Id_guru = guruId
	helper.ReadFromRequest(request, &requestUpdateguru)
	responseGuru := control_guru.Service_guru.Update(request.Context(), requestUpdateguru)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseGuru,
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control_guru *ControllerGuruImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	guruId := params.ByName("guruId")
	control_guru.Service_guru.Delete(request.Context(), guruId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control_guru *ControllerGuruImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	guruId := params.ByName("guruId")
	responseGuru := control_guru.Service_guru.FindById(request.Context(), guruId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseGuru,
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control_guru *ControllerGuruImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responseGuru := control_guru.Service_guru.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseGuru,
	}
	helper.WriteFromResponse(writer, &webResponse)
}
