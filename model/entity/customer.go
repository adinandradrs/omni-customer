package entity

import (
	"database/sql"
	"time"
)

type Customer struct {
	Id             uint   `json:"id" gorm:"primary_key"`
	Email          string `json:"email" gorm:"unique;size:350"`
	Fullname       string `json:"fullname"`
	Password       string `json:"password" gorm:"size:500"`
	Status         uint   `json:"status"`
	IsDeleted      bool   `json:"is_deleted"`
	ActivationId   string `json:"activation_id" gorm:"unique;size:300"`
	CreatedDate    time.Time
	UpdatedDate    sql.NullTime `json:"updated_date"`
	ActivationDate sql.NullTime `json:"activation_date"`
}
