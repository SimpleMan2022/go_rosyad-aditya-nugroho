package config

import (
	"fmt"
	model_fruits "fruits_project/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func LoadDb() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%s", ENV.DB_USERNAME, ENV.DB_PASSWORD,
		ENV.DB_URL, ENV.DB_NAME, "charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model_fruits.Nutritions{})
	db.AutoMigrate(&model_fruits.Fruit{})
	DB = db

}
