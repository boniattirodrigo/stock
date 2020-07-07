package models

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Ticker string  `gorm:"size:12;unique_index;not null"`
	Price  float64 `sql:"type:decimal(10,2);"`
}
