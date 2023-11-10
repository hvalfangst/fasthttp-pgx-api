package model

import "time"

type Car struct {
	ID           int64     `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	VIN          string    `json:"vin"`
	Color        string    `json:"color"`
	PurchaseDate time.Time `json:"purchase_date"`
	OwnerID      int       `json:"owner_id"`
	_            struct{}  `pg:"_schema:cars"`
}
