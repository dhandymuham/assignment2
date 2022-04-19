package controllers

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_ID   uint      `json:"customer_id"`
	Customer_Name string    `json:"name"`
	OrderedAt     time.Time `json:"ordered_at"`
	Items         []Items   `gorm:"foreignKey:Order_ID"`
}

type Items struct {
	gorm.Model
	Order_ID    uint   `json:"order_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
