package model

import "gorm.io/gorm"

type Purchaseorder struct {
    gorm.Model
    Name     string
    Email    string
    ItemlinesID uint // Foreign key referencing the Address
    Itemlines   Itemlines
}

// address.go
