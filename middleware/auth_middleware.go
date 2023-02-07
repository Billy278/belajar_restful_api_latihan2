package middleware

import (
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "MRGINZ" == request.Header.Get("PAS-API-KEY") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UnAuthorized",
		}
		helper.WriteFromResponse(writer, &webResponse)
	}
}
