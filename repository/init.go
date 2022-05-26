// init repository
/*
TO connect to mysql
   creat user table and video table
*/

package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() error {

	db, err := connectToDB()
	if err != nil {
		return err
	}

	if err := createUserTable(db); err != nil {
		return err
	}

	if err := createVideoTable(db); err != nil {
		return err
	}

	return nil
}

//----------------------------------------------------------------------------------------------------------------

func createUserTable(db *gorm.DB) error {

	// creat usertable by User struct
	if err := db.AutoMigrate(User{}); err != nil {
		return err
	}

	UserDB = db

	return nil
}

func createVideoTable(db *gorm.DB) error {

	// creat videotable by User struct
	if err := db.AutoMigrate(&Video{}); err != nil {
		return err
	}

	VideoDB = db

	return nil
}

//creat a dsn string to connect to mysql
func setDSN() string {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "douyin"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)

	return dsn
}

//connect to database and return a DB
func connectToDB() (*gorm.DB, error) {
	// connect to mysql by dsn
	dsn := setDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
