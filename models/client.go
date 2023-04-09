package models

import "gorm.io/gorm"

// Client
type Client struct {
	gorm.Model
	Name        string
	Description string
}
