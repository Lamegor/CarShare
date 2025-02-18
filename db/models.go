package db

import "time"

// Пользователь
type User struct {
	ID             int       `gorm:"primaryKey"`
	Login          string    `gorm:"unique,column:login"`
	Password       []byte    `gorm:"column:password"`
	FirstName      string    `gorm:"column:first_name"`
	SecondName     string    `gorm:"column:second_name"`
	LastName       string    `gorm:"column:last_name"`
	PassportData   string    `gorm:"column:passport_data"`
	IDLicense      string    `gorm:"column:id_license"`
	IsActive       bool      `gorm:"column:is_active"`
	Gender         string    `gorm:"column:gender"`
	Birthday       time.Time `gorm:"column:birthday"`
	ContactPhone   string    `gorm:"unique,column:contact_phone"`
	Email          string    `gorm:"unique,column:email"`
	Cookie         string    `gorm:"unique,column:cookie"`
	SessionExpires time.Time `gorm:"column:session_expires"`
}

// Модель машины
type Car struct {
	ID                int    `gorm:"primaryKey"`
	Manufacturer      string `gorm:"column:manufacturer"`
	Model             string `gorm:"column:model"`
	YearOfManufacture int    `gorm:"column:year_of_manufacture"`
	Condition         string `gorm:"column:condition"`
	IsActive          bool   `gorm:"column:is_active"`
	VIN               string `gorm:"column:vin"`
}
