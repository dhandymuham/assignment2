package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Order_ID    uint   `json:"order_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
