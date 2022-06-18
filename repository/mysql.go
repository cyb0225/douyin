// 连接数据库，创建相关表

package repository

import (
	"fmt"

	"github.com/2103561941/douyin/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() error {

	// 连接数据库
	err := connectToDB()
	if err != nil {
		return err
	}

	// 创建用户表
	if err := createUserTable(); err != nil {
		return err
	}

	// 创建视频表
	if err := createVideoTable(); err != nil {
		return err
	}

	// 创建关注表
	if err := createFollowTable(); err != nil {
		return err
	}

	// 创建点赞表
	if err := createLikeTable(); err != nil {
		return err
	}
	
	// 创建评论表
	if err := createCommentTable(); err != nil {
		return err
	}

	return nil
}

//----------------------------------------------------------------------------------------------------------------

// 创建关注表
func createUserTable() error {

	// creat usertable by User struct
	if err := Db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}

// 创建视频表
func createVideoTable() error {

	// creat videotable by User struct
	if err := Db.AutoMigrate(&Video{}); err != nil {
		return err
	}

	return nil
}

// 创建关注表
func createFollowTable() error {

	if err := Db.AutoMigrate(&Follow{}); err != nil {
		return err
	} // create follow table

	return nil
}

// 创建点赞表
func createLikeTable() error {

	if err := Db.AutoMigrate(&LikeTable{}); err != nil {
		return err
	} // create follow table

	return nil
}

// 创建评论表
func createCommentTable() error {

	if err := Db.AutoMigrate(&CommentTable{}); err != nil {
		return err
	} // create follow table

	return nil
}


// 创建mysql dsn字符串，用于连接mysql服务器
func setDSN() string {
	username := config.DBconf.Username
	password := config.DBconf.Password
	host := config.DBconf.Host
	port := config.DBconf.Port
	Dbname := config.DBconf.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)

	return dsn
}

// 连接数据库
func connectToDB() error {
	
	dsn := setDSN()
	var err error
	if Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}

	return nil
}
