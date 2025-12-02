package main

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type User struct {
	Name string
	Age  int64
}

func main() {
	user := ""
	pw := ""
	port := 0

	dsn := fmt.Sprintf("sqlserver://%s:%s@localhost:%d?database=master", user, pw, port)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err.Error())
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		print(err.Error())
		return
	}
	defer sqlDB.Close()
	dbname := "gorm"

	query := fmt.Sprintf("SELECT DB_ID('%s')", dbname)
	var dbID *int32
	result := db.Raw(query).Scan(&dbID)
	if result.Error != nil {
		err = db.Exec("CREATE DATABASE " + dbname).Error
		if err != nil {
			print(err.Error())
			return
		}
	} else if dbID == nil {
		err = db.Exec("CREATE DATABASE " + dbname).Error
		if err != nil {
			print(err.Error())
			return
		}
	}

	dsn2 := fmt.Sprintf("sqlserver:///%s:%s@localhost:%d?database=gorm", user, pw, port)
	db2, err := gorm.Open(sqlserver.Open(dsn2), &gorm.Config{})
	if err != nil {
		print(err.Error())
		return
	}
	sqlDB2, err := db2.DB()
	if err != nil {
		print(err.Error())
		return
	}
	defer sqlDB2.Close()

	err = db2.AutoMigrate(&User{})
	if err != nil {
		print(err.Error())
		return
	}
	err = db2.Create(User{Name: "a", Age: 10}).Error
	if err != nil {
		print(err.Error())
		return
	}
}
