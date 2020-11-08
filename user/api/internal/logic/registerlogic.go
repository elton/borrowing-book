package logic

import (
	"context"

	"github.com/elton/borrowing-book/user/api/internal/svc"
	"github.com/elton/borrowing-book/user/api/internal/types"
	"github.com/elton/borrowing-book/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) error {
	user, _ := l.svcCtx.UserModel.FindOneByName(req.Username)
	if user != nil {
		return errorDuplicateUsername
	}

	user, _ = l.svcCtx.UserModel.FindOneByMobile(req.Mobile)
	if user != nil {
		return errorDuplicateMobile
	}

	_, err := l.svcCtx.UserModel.Insert(model.User{
		Name:     req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   "男",
		Nickname: "admin",
	})

	return err
}
