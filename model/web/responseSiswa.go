package web

type ResponseSiswa struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Nis    string `json:"nis"`
	Alamat string `json:"alamat"`
}
