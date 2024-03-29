package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/utils"
	"github.com/go-water/water"
)

type LoginPostRequest struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginPostService struct {
	*water.ServerBase
}

func (srv *LoginPostService) Handle(ctx context.Context, req *LoginPostRequest) (interface{}, error) {
	auth, err := model.GetAuth(model.DbMap, req.User, req.Password)
	if err != nil {
		return nil, err
	}

	if auth == nil {
		return nil, errors.New("账号或密码错误")
	}

	token, err := water.SetAuthToken(req.User, "go-water", utils.RsaPrivateKeyPath, utils.AuthTimeout)
	if err != nil {
		return nil, errors.New("创建登陆失败")
	}

	return token, nil
}
