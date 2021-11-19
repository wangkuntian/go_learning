package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name string
	Age  uint8
	gorm.Model
}

func main() {
	dsn := "root:199412@tcp(127.0.0.1:3306)/moonflower?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	db.AutoMigrate(User{})
}
