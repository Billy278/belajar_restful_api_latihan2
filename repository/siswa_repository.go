package repository

import (
	"belajar-go-restful-api-latihan2/model/domain"
	"context"
	"database/sql"
)

type SiswaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa)
	FindByid(ctx context.Context, tx *sql.Tx, Siswaid int) (domain.Siswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Siswa
}
