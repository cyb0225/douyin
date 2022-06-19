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

// type RedisUser struct {
// 	Id            string
// 	Username      string
// 	Password      string
// 	FollowCount   string
// 	FollowerCount string
// }

func (*User) TableName() string {
	return "user"
}

// func (user *User) RedisConvert(redisUser *RedisUser) error {
// 	var err error

// 	user.Id, err = strconv.ParseUint(redisUser.Id, 10, 64)
// 	if err != nil {
// 		return errors.New("Redis conversion - userID error")
// 	}

// 	user.Username = redisUser.Username

// 	user.Password = redisUser.Password

// 	user.FollowCount, err = strconv.ParseUint(redisUser.FollowCount, 10, 64)
// 	if err != nil {
// 		return errors.New("Redis conversion - FollowCount error")
// 	}

// 	user.FollowerCount, err = strconv.ParseUint(redisUser.FollowerCount, 10, 64)
// 	if err != nil {
// 		return errors.New("Redis conversion - FollowerCount error")
// 	}
// 	return nil

// }

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
	log.Println("REDISTEST------------------")
	log.Println(user.Id)
	log.Println("REDISTEST------------------")
	redisResult, err := redis.Values(Client.Do("HGETALL", user.Id))
	if err != nil {
		return err
	}

	/*
			这里需要将redis缓存的内容解析。如果长度为0则证明没有东西，需要去SQL拿。
			如果不为0，应将数据解析然后返回。
			redis存的数据长这样：
			1) "username"
			2) "cyb123"
			3) "password"
			4) "8bfc6a7a8d1355bddae62ba8d898ac57"
			5) "follow_count"
			6) "5"
			7) "follower_count"
			8) "0"


		理论上的顺序：
		服务器刚启动，redis为空，mysql有数据。
		一开始肯定redis找不到，去mysql找。然后mysql返回值。
		在三个用了SelectByUserId函数的地方，下面紧接着用了RewriteToRedis函数，把获得的数据写回redis

		然后如果发生数据更新情况，针对user表是follow和follower数量的更新 也就是函数UpdateFollowCount和UpdateFollowerCount
		数量更新对应的函数已经做了redis数据的更新操作。
		理论上不会有问题。但是还应该多检查一下。

		注意convert是必要的。因为redis只能存string。迟早要转一遍。但我不知道有没有更牛逼的方法。


	*/
	if len(redisResult) > 0 { //如果redis有数据, 直接存到user里
		if err := redis.ScanStruct(redisResult, user); err != nil { //原始redis返回值写入redisconvert结构 全部为string类型
			fmt.Println(err)
		}
		log.Println("Redis Found UserData")
		// var result *User
		// err := result.RedisConvert(redisconvert) //转换格式并写入result返回值
		// if err != nil {
		// return errors.New("redis conversion error")
		// }
		//-------------------------------

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
