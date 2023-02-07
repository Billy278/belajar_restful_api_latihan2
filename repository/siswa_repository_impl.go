package repository

import (
	"belajar-go-restful-api-latihan2/helper"
	"belajar-go-restful-api-latihan2/model/domain"
	"context"
	"database/sql"
	"errors"
)

type SiswaRepositoryImpl struct {
}

func NewSiswaRepositoryImpl() SiswaRepository {
	return &SiswaRepositoryImpl{}
}

func (siswa_repository *SiswaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	sql := "Insert Into Siswa(name,nis,alamat) Values(?,?,?)"
	result, err := tx.ExecContext(ctx, sql, siswa.Name, siswa.Nis, siswa.Alamat)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	siswa.Id = int(id)
	return siswa

}

func (siswa_repository *SiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	sql := "Update siswa set name=?,nis=?,alamat=? Where id=?"
	_, err := tx.ExecContext(ctx, sql, siswa.Name, siswa.Nis, siswa.Alamat, siswa.Id)
	helper.PanicIfError(err)
	return siswa

}
func (siswa_repository *SiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) {
	sql := "Delete From siswa Where id=?"
	_, err := tx.ExecContext(ctx, sql, siswa.Id)
	helper.PanicIfError(err)
}

func (siswa_repository *SiswaRepositoryImpl) FindByid(ctx context.Context, tx *sql.Tx, Siswaid int) (domain.Siswa, error) {
	sql := "Select id,name, nis,alamat From siswa Where id=?"
	rows, err := tx.QueryContext(ctx, sql, Siswaid)
	helper.PanicIfError(err)
	defer rows.Close()
	siswa := domain.Siswa{}
	if rows.Next() {
		err := rows.Scan(&siswa.Id, &siswa.Name, &siswa.Nis, &siswa.Alamat)
		helper.PanicIfError(err)

		return siswa, nil
	} else {
		return siswa, errors.New("SISWA IS NOT FOUND")
	}

}

func (siswa_repository *SiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Siswa {
	sql := "Select id,name,nis,alamat from siswa"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var siswaAll []domain.Siswa
	for rows.Next() {
		siswa := domain.Siswa{}
		err := rows.Scan(&siswa.Id, &siswa.Name, &siswa.Nis, &siswa.Alamat)
		helper.PanicIfError(err)
		siswaAll = append(siswaAll, siswa)
	}

	return siswaAll
}
