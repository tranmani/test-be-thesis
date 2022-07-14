package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID int64 `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
}

func GetAll() ([]User, error) {
	var users []User

	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		return users, err
	}

	db.Table("test").Select("*").Scan(&users)

	var result []User

	for i := 0; i < 10; i++ {
    result = append(result, users...)
}


	return result, nil
}