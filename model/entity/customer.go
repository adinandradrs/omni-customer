package entity

import (
	"database/sql"
	"time"
)

type Customer struct {
	Id             uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Email          sql.NullString `json:"email" gorm:"unique;size:350"`
	Fullname       sql.NullString `json:"fullname"`
	Password       sql.NullString `json:"password" gorm:"size:500"`
	PhoneNo        sql.NullString `json:"phone_no" gorm:"unique;size:15"`
	Address        sql.NullString `json:"address" gorm:"size:350"`
	Status         uint           `json:"status"`
	IsDeleted      bool           `json:"is_deleted"`
	ActivationId   sql.NullString `json:"activation_id" gorm:"unique;size:300"`
	CreatedDate    time.Time
	UpdatedDate    sql.NullTime `json:"updated_date"`
	ActivationDate sql.NullTime `json:"activation_date"`
}
