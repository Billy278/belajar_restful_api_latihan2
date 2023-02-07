package service

import (
	"belajar-go-restful-api-latihan2/model/web"
	"context"
)

type ServiceSiswaRepository interface {
	Create(ctx context.Context, request web.RequestCreateSiswa) web.ResponseSiswa
	Update(ctx context.Context, request web.RequestUpdateSiswa) web.ResponseSiswa
	Delete(ctx context.Context, siswaId int)
	FindById(ctx context.Context, siswaId int) web.ResponseSiswa
	FindAll(ctx context.Context) []web.ResponseSiswa
}
