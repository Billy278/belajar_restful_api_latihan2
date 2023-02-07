package web

import "time"

type RequestUpdateGuru struct {
	Id_guru   string    `validate:"required" json:"id_guru"`
	Name      string    `validate:"required" json:"name"`
	Birth_day time.Time `validate:"required" json:"birth_day"`
	Nig       string    `validate:"required" json:"nig"`
	Status    *bool     `validate:"required" json:"status"`
}
