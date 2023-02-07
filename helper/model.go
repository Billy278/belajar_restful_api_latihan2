package helper

import (
	"belajar-go-restful-api-latihan2/model/domain"
	"belajar-go-restful-api-latihan2/model/web"
)

func ToSiswaResponse(siswa domain.Siswa) web.ResponseSiswa {
	return web.ResponseSiswa{
		Id:     siswa.Id,
		Name:   siswa.Name,
		Nis:    siswa.Nis,
		Alamat: siswa.Alamat,
	}

}

func ToGuruResponse(guru domain.Guru) web.ResponseGuru {
	return web.ResponseGuru{
		Id_guru:   guru.Id_guru,
		Name:      guru.Name,
		Birth_day: guru.Birth_day,
		Nig:       guru.Nig,
		Status:    guru.Status,
	}
}

func ToSiswaResponses(Allsiswa []domain.Siswa) []web.ResponseSiswa {
	var webResponses []web.ResponseSiswa
	for _, siswa := range Allsiswa {
		webResponses = append(webResponses, ToSiswaResponse(siswa))
	}
	return webResponses

}

func ToGuruResponses(Allguru []domain.Guru) []web.ResponseGuru {
	var webResponses []web.ResponseGuru
	for _, guru := range Allguru {
		webResponses = append(webResponses, ToGuruResponse(guru))
	}
	return webResponses

}
