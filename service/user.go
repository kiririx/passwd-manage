package service

import (
	"github.com/kiririx/krutils/algox"
	"github.com/kiririx/passwd-manage/conf"
	"github.com/kiririx/passwd-manage/module/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type userService struct {
}

func (u *userService) QueryByUsername(username string) (*entity.User, error) {
	var user entity.User
	tx := conf.Sqlx.Where("username = ?", username).Take(&user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userService) Save(user *entity.User) bool {
	result := conf.Sqlx.Create(user)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return false
	}
	return true
}

func (u *userService) Query(e *entity.User) any {
	first := conf.Sqlx.First(e)
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	if first.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return e
}

func (u *userService) Register(username string, password string) (*entity.User, error) {
	user, err := User.QueryByUsername(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户已注册")
	}
	userEntity := &entity.User{
		Username: username,
		Password: algox.MD5(password),
	}
	tx := conf.Sqlx.Save(userEntity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, errors.New("注册失败")
	}
	return userEntity, nil
}

func (u *userService) Login(username string, password string) (string, error) {
	var user entity.User
	if err := conf.Sqlx.Where("username = ? and password = ?", username, algox.MD5(password)).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户不存在或密码错误")
		}
		return "", err
	}
	return BuildToken(user.Id, user.Username)
}
