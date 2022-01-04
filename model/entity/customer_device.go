package entity

import (
	"time"
)

type CustomerDevice struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	DeviceName  string `json:"device_name"`
	OS          string `json:"operating_system"`
	IsDeleted   bool   `json:"is_deleted"`
	DeviceId    string `json:"device_id" gorm:"unique;size:300"`
	CreatedDate time.Time
}
