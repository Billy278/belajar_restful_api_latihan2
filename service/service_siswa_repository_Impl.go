package service

import (
	"belajar-go-restful-api-latihan2/exception"
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/domain"
	"belajar-go-restful-api-latihan2/model/web"
	"belajar-go-restful-api-latihan2/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ServiceSiswaRepositoryImpl struct {
	DB              *sql.DB
	validate        *validator.Validate
	repositorySiswa repository.SiswaRepository
}

func NewServiceSiswaRepositoryImpl(db *sql.DB, validate *validator.Validate, repository repository.SiswaRepository) ServiceSiswaRepository {
	return &ServiceSiswaRepositoryImpl{
		DB:              db,
		validate:        validate,
		repositorySiswa: repository,
	}
}

func (service_siswa *ServiceSiswaRepositoryImpl) Create(ctx context.Context, request web.RequestCreateSiswa) web.ResponseSiswa {
	err := service_siswa.validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service_siswa.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa := domain.Siswa{
		Name:   request.Name,
		Nis:    request.Nis,
		Alamat: request.Alamat,
	}
	siswa = service_siswa.repositorySiswa.Create(ctx, tx, siswa)
	siswaResponse := web.ResponseSiswa{
		Id:     siswa.Id,
		Name:   siswa.Name,
		Nis:    siswa.Nis,
		Alamat: siswa.Alamat,
	}
	return siswaResponse
}

func (service_siswa *ServiceSiswaRepositoryImpl) Update(ctx context.Context, request web.RequestUpdateSiswa) web.ResponseSiswa {
	err := service_siswa.validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service_siswa.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa, err := service_siswa.repositorySiswa.FindByid(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//helper.PanicIfError(err)
	siswa.Name = request.Name
	siswa.Nis = request.Nis
	siswa.Alamat = request.Alamat
	siswa = service_siswa.repositorySiswa.Update(ctx, tx, siswa)
	return helper.ToSiswaResponse(siswa)
}

func (service_siswa *ServiceSiswaRepositoryImpl) Delete(ctx context.Context, siswaId int) {
	tx, err := service_siswa.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa, err := service_siswa.repositorySiswa.FindByid(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//helper.PanicIfError(err)
	service_siswa.repositorySiswa.Delete(ctx, tx, siswa)

}

func (service_siswa *ServiceSiswaRepositoryImpl) FindById(ctx context.Context, siswaId int) web.ResponseSiswa {
	tx, err := service_siswa.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa, err := service_siswa.repositorySiswa.FindByid(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//	helper.PanicIfError(err)
	return helper.ToSiswaResponse(siswa)
}

func (service_siswa *ServiceSiswaRepositoryImpl) FindAll(ctx context.Context) []web.ResponseSiswa {
	tx, err := service_siswa.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa := service_siswa.repositorySiswa.FindAll(ctx, tx)
	return helper.ToSiswaResponses(siswa)

}
