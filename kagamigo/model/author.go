package model

import "gorm.io/gorm"

type Author struct {
    gorm.Model
    Name string `gorm:"unique"`
}

type Post struct {
    gorm.Model
    Title string
    AuthorID uint `gorm:"foreignKey:AuthorID"`
    Author Author `gorm:"foreignKey:ID"`
}

type PostUser struct {
    gorm.Model
    PostID uint `gorm:"foreignKey:PostID"`
    AuthorID uint `gorm:"foreignKey:AuthorID"`
}