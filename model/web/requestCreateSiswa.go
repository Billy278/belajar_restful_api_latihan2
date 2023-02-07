package web

type RequestCreateSiswa struct {
	Name   string `validate:"required" json:"name"`
	Nis    string `validate:"required" json:"nis"`
	Alamat string `validate:"required" json:"alamat"`
}
