package model

import "gorm.io/gorm"

type Book struct {
    gorm.Model
    Title    string
    AuthorID uint // Foreign key referencing the Author
    Author   Author
}