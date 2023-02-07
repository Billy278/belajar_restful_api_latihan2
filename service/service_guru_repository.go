package service

import (
	"belajar-go-restful-api-latihan2/model/web"
	"context"
)

type ServiceGuruRepository interface {
	Create(ctx context.Context, request web.RequestCreateGuru) web.ResponseGuru
	Update(ctx context.Context, request web.RequestUpdateGuru) web.ResponseGuru
	Delete(ctx context.Context, guruId string)
	FindById(ctx context.Context, guruId string) web.ResponseGuru
	FindAll(ctx context.Context) []web.ResponseGuru
}
