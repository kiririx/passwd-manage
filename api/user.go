package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/passwd-manage/module/req"
	"github.com/kiririx/passwd-manage/service"
)

type UserApi struct {
}

// Register 用户注册
func (u *UserApi) Register(c *gin.Context, param *req.Register) (any, error) {
	_, err := service.User.Register(param.Username, param.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Login 用户登陆
func (u *UserApi) Login(c *gin.Context, param *req.Login) (any, error) {
	token, err := service.User.Login(param.Username, param.Password)
	if err != nil {
		return nil, err
	}
	return map[string]string{"token": token}, nil
}
