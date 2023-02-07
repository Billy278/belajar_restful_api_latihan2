package repository

import (
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/domain"
	"context"
	"database/sql"
	"errors"
)

type GuruRepositoryImpl struct {
}

func NewGuruRepositoryImpl() GuruRepository {
	return &GuruRepositoryImpl{}
}

func (guru_repository *GuruRepositoryImpl) Create(ctx context.Context, db *sql.DB, guru domain.Guru) domain.Guru {
	sql := "Insert Into guru(id_guru,name,birth_day,nig,status) Values(?,?,?,?,?)"
	_, err := db.ExecContext(ctx, sql, guru.Id_guru, guru.Name, guru.Birth_day, guru.Nig, guru.Status)
	helper.PanicIfError(err)
	return guru
}

func (guru_repository *GuruRepositoryImpl) Update(ctx context.Context, db *sql.DB, guru domain.Guru) domain.Guru {
	sql := "UPDATE guru set name=?,birth_day=?,nig=?,status=? WHERE id_guru=?"
	_, err := db.ExecContext(ctx, sql, guru.Name, guru.Birth_day, guru.Nig, guru.Status, guru.Id_guru)
	helper.PanicIfError(err)
	return guru
}

func (guru_repository *GuruRepositoryImpl) Delete(ctx context.Context, db *sql.DB, guru domain.Guru) {
	sql := "Delete from guru where id_guru=?"
	_, err := db.ExecContext(ctx, sql, guru.Id_guru)
	helper.PanicIfError(err)

}

func (guru_repository *GuruRepositoryImpl) FindByid(ctx context.Context, db *sql.DB, guruId string) (domain.Guru, error) {
	sql := "Select id_guru,name,birth_day,nig,status from guru Where id_guru=?"
	rows, err := db.QueryContext(ctx, sql, guruId)
	helper.PanicIfError(err)
	defer rows.Close()
	guru := domain.Guru{}
	if rows.Next() {
		err := rows.Scan(&guru.Id_guru, &guru.Name, &guru.Birth_day, &guru.Nig, &guru.Status)
		helper.PanicIfError(err)
		return guru, nil
	} else {
		return guru, errors.New("NOT FOUND")
	}
}

func (guru_repository *GuruRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Guru {
	sql := "Select id_guru,name,birth_day,nig,status from guru"
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var allguru []domain.Guru
	for rows.Next() {
		guru := domain.Guru{}
		err := rows.Scan(&guru.Id_guru, &guru.Name, &guru.Birth_day, &guru.Nig, &guru.Status)
		helper.PanicIfError(err)
		allguru = append(allguru, guru)
	}
	return allguru
}
