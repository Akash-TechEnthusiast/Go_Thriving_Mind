package model

import "gorm.io/gorm"

type Itemlines struct {
    gorm.Model
    Street  string
    City    string
    Country string
}