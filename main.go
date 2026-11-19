package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Temp struct {
	gorm.Model
	Test string
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbname   = "projects"
	port     = "5432"
	schema   = "rbac"
	sslmode  = "disable"
	timeZone = "Asia/Tehran"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbname,
		port,
		schema,
		sslmode,
		timeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error; err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&Temp{}); err != nil {
		log.Fatal(err)
	}

	//create
	db.Create(&Temp{Test: "test string 1"})

	//read
	var temp Temp;
	db.First(&temp , "test = ?" , "test string 1")
	fmt.Printf("%#v\n" , temp)

	//update
	db.Model(&temp).Update("test" , "test string 2")
	db.First(&temp , "test = ?" , "test string 2")
	fmt.Printf("%#v\n" , temp)

	//delete(soft)
	db.Delete(&temp)
	db.First(&temp , "test = ?" , "test string 2")
	fmt.Printf("%#v\n" , temp)

	//delete(hard)
	db.Unscoped().Delete(&temp)
	db.First(&temp , "test = ?" , "test string 2")
	fmt.Printf("%#v\n" , temp)
}
