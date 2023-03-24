package entities

import "time"

type Transaksi struct {
	Id           int
	UserId       int
	Tanggal      time.Time
	CustomerId   int
	CustomerName string
	PhoneNum     string
	AddressCus   string
	EmplloyeName string
}
