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

type ServiceGuruRepositoryImpl struct {
	DB             *sql.DB
	repositoryGuru repository.GuruRepository
	validate       *validator.Validate
}

func NewServiceGuruRepositoryImpl(db *sql.DB, repository repository.GuruRepository, validate *validator.Validate) ServiceGuruRepository {
	return &ServiceGuruRepositoryImpl{
		DB:             db,
		repositoryGuru: repository,
		validate:       validate,
	}
}

func (service_guru *ServiceGuruRepositoryImpl) Create(ctx context.Context, request web.RequestCreateGuru) web.ResponseGuru {
	err := service_guru.validate.Struct(request)
	helper.PanicIfError(err)
	db := service_guru.DB
	//defer db.Close()
	guru := domain.Guru{
		Id_guru:   request.Id_guru,
		Name:      request.Name,
		Birth_day: request.Birth_day,
		Nig:       request.Nig,
		Status:    *request.Status,
	}
	guru = service_guru.repositoryGuru.Create(ctx, db, guru)
	return helper.ToGuruResponse(guru)
}

func (service_guru *ServiceGuruRepositoryImpl) Update(ctx context.Context, request web.RequestUpdateGuru) web.ResponseGuru {
	err := service_guru.validate.Struct(request)
	helper.PanicIfError(err)
	db := service_guru.DB
	//defer db.Close()
	guru, err := service_guru.repositoryGuru.FindByid(ctx, db, request.Id_guru)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//helper.PanicIfError(err)
	guru.Name = request.Name
	guru.Birth_day = request.Birth_day
	guru.Nig = request.Nig
	guru.Status = *request.Status
	guru = service_guru.repositoryGuru.Update(ctx, db, guru)
	return helper.ToGuruResponse(guru)

}

func (service_guru *ServiceGuruRepositoryImpl) Delete(ctx context.Context, guruId string) {
	db := service_guru.DB
	guru, err := service_guru.repositoryGuru.FindByid(ctx, db, guruId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//helper.PanicIfError(err)
	service_guru.repositoryGuru.Delete(ctx, db, guru)
}

func (service_guru *ServiceGuruRepositoryImpl) FindById(ctx context.Context, guruId string) web.ResponseGuru {
	db := service_guru.DB
	//defer db.Close()
	guru, err := service_guru.repositoryGuru.FindByid(ctx, db, guruId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//helper.PanicIfError(err)
	return helper.ToGuruResponse(guru)
}

func (service_guru *ServiceGuruRepositoryImpl) FindAll(ctx context.Context) []web.ResponseGuru {
	db := service_guru.DB
	//defer db.Close()
	guru := service_guru.repositoryGuru.FindAll(ctx, db)
	return helper.ToGuruResponses(guru)
}
