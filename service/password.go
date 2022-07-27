package service

import (
	"errors"
	"fmt"
	"github.com/kiririx/passwd-manage/conf"
	"github.com/kiririx/passwd-manage/module/entity"
	"gorm.io/gorm"
	"log"
)

type passwordService struct {
}

func (*passwordService) Save(passwd *entity.Password, userId uint) bool {
	passwd.UserId = userId
	if passwd.IsEmpty() {
		conf.Sqlx.Save(passwd)
	} else {
		conf.Sqlx.Updates(*passwd)
	}
	return true
}
func (*passwordService) QueryPasswordById(id uint) (passwd entity.Password) {
	err := conf.Sqlx.Where("id = ?", id).Take(&passwd).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// todo format for logging
		log.Println(err)
	}
	return
}

func (*passwordService) QueryPasswordListByUserId(userId uint, pageNum uint, pageSize uint) (passwords []entity.Password, total int64) {
	err := conf.Sqlx.Where("user_id = ?", userId).Limit(int(pageSize)).Offset(int(pageNum * pageSize)).Order("createTime desc").Find(&passwords).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// todo format for logging
		log.Println(err)
		return
	}
	err = conf.Sqlx.Model(&entity.Password{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// todo format for logging
		log.Println(err)
		return
	}
	return
}

func (*passwordService) FilterPasswordListByUserId(userId uint, topic string, pageNum uint, pageSize uint) (passwords []entity.Password, total int64) {
	err := conf.Sqlx.Where("user_id = ? and topic like ?", userId, fmt.Sprintf("%%%v%%", topic)).Limit(int(pageSize)).Offset(int(pageNum * pageSize)).Order("createTime desc").Find(&passwords).Error
	if err != nil {
		// todo format for logging
		log.Println(err)
	}
	err = conf.Sqlx.Model(&entity.Password{}).Where("user_id = ? and topic like ?", userId, fmt.Sprintf("%%%v%%", topic)).Count(&total).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// todo format for logging
		log.Println(err)
	}
	return
}

func (*passwordService) DeleteById(id uint) {
	err := conf.Sqlx.Where("id = ?", id).Delete(&entity.Password{}).Error
	if err != nil {
		// todo format for logging
		log.Println(err)
	}
	return
}
