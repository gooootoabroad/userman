package logic

import (
	"context"
	"fmt"

	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/model"
	"userman/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// todo: add your logic here and delete this line
	// 检查角色是否正确
	var role model.Role
	role, err = role.ToRole(req.Roles)
	if err != nil {
		l.Errorf("check user %s role %s failed, err: %s", req.Username, req.Roles, err)
		return nil, fmt.Errorf("user role is illegal")
	}

	user := model.UserInfo{
		Username: req.Username,
		NickName: req.NickName,
		Password: req.Password,
		Roles:    role,
	}

	l.Infof("start create user %s", req.Username)
	err = utils.CreateUser(l.ctx, user)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("create user %s successed", req.Username)
	resp = &types.CreateUserResp{
		Message: msg,
	}

	return resp, nil
}
