package models

type User struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    PhoneNumber string `json:"phone_number" gorm:"unique;not null"`
    Password    string `json:"password" gorm:"not null"`
}