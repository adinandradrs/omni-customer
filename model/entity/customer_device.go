package entity

import (
	"database/sql"
	"time"
)

type CustomerDevice struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	DeviceName  sql.NullString `json:"device_name"`
	OS          sql.NullString `json:"operating_system"`
	IsDeleted   bool           `json:"is_deleted"`
	DeviceId    sql.NullString `json:"device_id" gorm:"unique;size:300"`
	CreatedDate time.Time
}
