package models

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Ticker string  `gorm:"size:10;unique;not null"`
	Price  float64 `sql:"type:decimal(10,2);"`
}
