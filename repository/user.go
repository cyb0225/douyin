// store userinfos and related CRUD interface

package repository
import (
	"errors"

	"gorm.io/gorm"
)

var (
	UserDB *gorm.DB
)

// user mesages in database table
type User struct {
	ID       uint64 `gorm:"colomn:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Follow   Follow
}

// User relational tables
type Follow struct {
	ID uint64 `gorm:"colomn:id"`
}

type Follower struct {
	ID   uint64 `gorm:"colomn:id"`
	User User
}

type Publish struct {
	ID   uint64 `gorm:"colomn:id"`
	User User
}

type Favourite struct {
	ID   uint64 `gorm:"colomn:id"`
	User User
}

func (user *User) Insert() error {

	//insert error
	if err := UserDB.Table(user.TableName()).Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase error")
	}

	return nil
}

// query user record by username
func (user *User) SelectByUsername() error {

	result := UserDB.Where("username = ?", user.Username).First(user)

	// can't find record use the username
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}

	return nil
}

func (*User) TableName() string {
	return "user"
}
