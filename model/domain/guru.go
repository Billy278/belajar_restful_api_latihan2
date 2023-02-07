package domain

import (
	"time"
)

type Guru struct {
	Id_guru   string
	Name      string
	Birth_day time.Time
	Nig       string
	Status    bool
}
