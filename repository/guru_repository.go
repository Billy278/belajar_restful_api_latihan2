package repository

import (
	"belajar-go-restful-api-latihan2/model/domain"
	"context"
	"database/sql"
)

type GuruRepository interface {
	Create(ctx context.Context, db *sql.DB, guru domain.Guru) domain.Guru
	Update(ctx context.Context, db *sql.DB, guru domain.Guru) domain.Guru
	Delete(ctx context.Context, db *sql.DB, guru domain.Guru)
	FindByid(ctx context.Context, db *sql.DB, guruId string) (domain.Guru, error)
	FindAll(ctx context.Context, db *sql.DB) []domain.Guru
}
