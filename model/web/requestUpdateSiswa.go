package web

type RequestUpdateSiswa struct {
	Id     int    `validate:"required" json:"id"`
	Name   string `validate:"required" json:"name"`
	Nis    string `validate:"required" json:"nis"`
	Alamat string `validate:"required" json:"alamat"`
}
