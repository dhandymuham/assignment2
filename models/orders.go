package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Customer_ID   uint      `json:"customer_id"`
	Customer_Name string    `json:"name"`
	OrderedAt     time.Time `json:"ordered_at"`
	Items         []Items   `gorm:"foreignKey:Order_ID"`
}
