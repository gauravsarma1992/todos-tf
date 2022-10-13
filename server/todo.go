package server

import (
	"gorm.io/gorm"
)

type (
	TodoItem struct {
		gorm.Model
		Item   string `json:"item"`
		Status string `json:"status"`
	}
)
