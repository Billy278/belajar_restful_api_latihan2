package web

import "time"

type ResponseGuru struct {
	Id_guru   string    `json:"id_guru"`
	Name      string    `json:"name"`
	Birth_day time.Time `json:"birth_day"`
	Nig       string    `json:"nig"`
	Status    bool      `json:"status"`
}
