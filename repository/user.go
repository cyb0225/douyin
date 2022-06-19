// 用户信息

package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

type User struct {
	Id            uint64 `gorm:"column:id"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"follow_count"`
	FollowerCount uint64 `gorm:"follower_count"`
}


func (*User) TableName() string {
	return "user"
}


func (user *User) RewriteToRedis() error { //数据写回redis
	if _, err := Client.Do("HSET", user.Id, "Username", user.Username, "Password", user.Password, "FollowCount", user.FollowCount, "FollowerCount", user.FollowerCount); err != nil { //redis 写入
		return errors.New("REDIS----Insert to UserDatabase error, roll backed")
	}
	return nil
}

// 插入数据
func (user *User) Insert() error {
	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	//insert error
	if err := tx.Table(user.TableName()).Create(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("Insert to UserDatabase error, roll backed")
	}
	log.Println(user.Id, user.Username, user.Password, user.FollowCount, user.FollowerCount)

	if err := user.RewriteToRedis(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	mutex.Unlock()

	return nil

}

// select user record by username
func (user *User) SelectByUsername() error {

	result := Db.Table(user.TableName()).Where("username = ?", user.Username).First(user)

	// not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// select user record by user_id

func (user *User) SelectByUserId() error {

	redisResult, err := redis.Values(Client.Do("HGETALL", user.Id))
	if err != nil {
		return err
	}

	if len(redisResult) > 0 { //如果redis有数据, 直接存到user里
		if err := redis.ScanStruct(redisResult, user); err != nil { //原始redis返回值写入redisconvert结构 全部为string类型
			fmt.Println(err)
		}
		log.Println("Redis Found UserData")


	} else {
		log.Println("Redis NOT Found UserData")
		result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user)

		// not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New(result.Error.Error())
		}

		// 存到redis里
		//.....
		if err := user.RewriteToRedis(); err != nil {
			return errors.New("redis insert error")
		}
	}

	log.Println("select finish!!!!")

	return nil
}

// updata user follow_count by add n
func (user *User) UpdateFollowCount(n int) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follow_count", int(user.FollowCount)+n).Error; err != nil {
		tx.Rollback()
		return errors.New("update follow count error, roll backed")

	}
	if err := user.RewriteToRedis(); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	mutex.Unlock()

	return nil

}

// updata user follower_count by add n
func (user *User) UpdateFollowerCount(n int) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follower_count", int(user.FollowerCount)+n).Error; err != nil {
		tx.Rollback()
		return errors.New("update follower count error, roll backed")

	}

	if err := user.RewriteToRedis(); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	mutex.Unlock()
	return nil

}
