package model

import "gorm.io/gorm"

type Finances struct {
	gorm.Model
	NetWorth    float64
	Assets      float64
	Liabilities float64
}
