// 初始化仓库，连接mysql，提供对外接口

package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化用户仓库，与数据库进行连接
func InitUserRep() error {
	// 数据库连接
	dsn := setDSN()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 根据指定的类，创建表（或增加没有的列）
	database.AutoMigrate(&UserInfo{})

	// 赋值给db，供user增删改查使用
	db = database
	return nil
}

// 配置dsn
func setDSN() string {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "douyinUser"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	return dsn
}
