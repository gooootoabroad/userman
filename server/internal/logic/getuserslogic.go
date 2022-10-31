package logic

import (
	"context"
	"net/http"

	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers() (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line
	users, err := utils.GetAllUsers(l.ctx)
	if err != nil {
		l.Logger.Errorf("get users failed, err: %v", err)
		return nil, err
	}

	data := []types.UserInfo{}
	for index := range users {
		user := users[index]
		role, _ := user.Roles.String()
		lock := ""
		if user.LockAt != nil {
			lock = user.LockAt.String()
		}
		userInfo := types.UserInfo{
			UUID:      user.UUID.String(),
			Username:  user.Username,
			NickName:  user.NickName,
			Roles:     role,
			LockAt:    lock,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.CreatedAt.Unix(),
		}

		data = append(data, userInfo)
	}

	resp = &types.GetUserResp{
		Code:    http.StatusOK,
		Message: "",
		Data:    data,
	}

	return resp, nil
}
