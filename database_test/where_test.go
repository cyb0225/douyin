package databasetest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/2103561941/douyin/user/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setDSN() string {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "douyinUser"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	return dsn
}

// 测试能不能搜到数据库里面的数据
func TestDatabaseWhere(t *testing.T) {
	dsn := setDSN()

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log(err)
	}

	user := repository.UserInfo{
		Username: "22222",
	}

	db.AutoMigrate()

	result := db.Where("username = ?", user.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		t.Log("record not found")
		t.Log("11111111111111111111111")
		return
	}

	t.Log(user.Username, user.Password)

	// if err := repository.InitUserRep(); err != nil {
	// 	t.Log(err)
	// 	return
	// }

	// user := repository.UserInfo{
	// 	Username: "cyb123",
	// }

	// // 找不到
	// if err := user.QueryByUsername(); err != nil {
	// 	t.Log("1111111111111111111")
	// 	t.Log(err)
	// 	return
	// }

	// // 找到了
	// t.Log(user.Username, user.Password)

}
