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

var Db *gorm.DB

func Init() error {

	err := connectToDB()
	if err != nil {
		return err
	}

	if err := createUserTable(); err != nil {
		return err
	}

	if err := createVideoTable(); err != nil {
		return err
	}

	if err := createFollowTable(); err != nil {
		return err
	}

	if err := createLikeTable(); err != nil {
		return err
	}
	if err := createCommentTable(); err != nil {
		return err
	}

	return nil
}

//----------------------------------------------------------------------------------------------------------------

func createUserTable() error {

	// creat usertable by User struct
	if err := Db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}

func createVideoTable() error {

	// creat videotable by User struct
	if err := Db.AutoMigrate(&Video{}); err != nil {
		return err
	}

	return nil
}

func createFollowTable() error {

	if err := Db.AutoMigrate(&Follow{}); err != nil {
		return err
	} // create follow table

	return nil
}

func createLikeTable() error {

	if err := Db.AutoMigrate(&LikeTable{}); err != nil {
		return err
	} // create follow table

	return nil
}

func createCommentTable() error {

	if err := Db.AutoMigrate(&CommentTable{}); err != nil {
		return err
	} // create follow table

	return nil
}

//creat a dsn string to connect to mysql
func setDSN() string {
	username := "root"
	password := "070831"
	host := "127.0.0.1"
	port := 3306
	Dbname := "testsql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)

	return dsn
}

//connect to database and return a DB
func connectToDB() error {
	// connect to mysql by dsn
	dsn := setDSN()
	var err error
	if Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}

	return nil
}
