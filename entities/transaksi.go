package entities

import "time"

type Transaksi struct {
	Id         int
	UserId     int
	Tanggal    time.Time
	CustomerId int
}
