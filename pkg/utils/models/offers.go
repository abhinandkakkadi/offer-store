package models

import "time"

type OfferCompany struct {
	OfferID            uint      `gorm:"primary_key"`
	ClientID           uint      `gorm:"not null"`
	Country            string    `gorm:"not null"`
	Image              string    `gorm:"not null"`
	ImageWidth         int       `gorm:"not null;default:0"`
	ImageHeight        int       `gorm:"not null;default:0"`
	TextLocale         string    `gorm:"type:tinytext;not null"`
	ValidityTextLocale string    `gorm:"type:tinytext;not null"`
	Position           int       `gorm:"not null"`
	ValidFrom          time.Time `gorm:"not null;default:'1970-01-01 00:00:00'"`
	ShowFrom           time.Time `gorm:"default:'1970-01-01 00:00:00'"`
	ValidTo            time.Time `gorm:"not null;default:'2100-01-01 00:00:00'"`
	Flag               uint      `gorm:"not null;default:1"`
	PageCount          uint      `gorm:"not null;default:0"`
	StoreURL           string    `gorm:"default:''"`
	StoreURLTitle      string    `gorm:"not null"`
	OfferHome          int       `gorm:"not null"`
}
