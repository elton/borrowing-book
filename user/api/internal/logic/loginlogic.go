package logic

import (
	"context"

	"github.com/elton/borrowing-book/user/api/internal/svc"
	"github.com/elton/borrowing-book/user/api/internal/types"
	"github.com/elton/borrowing-book/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.UserReply, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	switch err {
	case nil:
		if user.Password != req.Password {
			return nil, errorIncorrectPassword
		}
		return &types.UserReply{
			Id:       user.Id,
			Username: user.Name,
			Mobile:   user.Mobile,
			Nickname: user.Nickname,
			Gender:   user.Gender,
		}, nil
	case model.ErrNotFound:
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}
}
