package logic

import (
	"context"
	"fmt"

	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/model"
	"userman/server/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line
	if req.Username == "" && req.UUID == "" {
		return nil, fmt.Errorf("need username or uuid")
	}

	user := &model.UserInfo{}
	if req.Username != "" {
		user, err = utils.GetUserByUsername(l.ctx, req.Username)
		if err != nil {
			l.Errorf("get user info by username %s failed, err: %s", req.Username, err)
			return nil, err
		}
	}

	if req.UUID != "" {
		userUUID, err := uuid.Parse(req.UUID)
		if err != nil {
			l.Errorf("change uuid %s failed, err: %s", req.UUID, err)
			return nil, err
		}

		uuidUser, err := utils.GetUserByUUID(l.ctx, userUUID)
		if err != nil {
			l.Errorf("get user info by uuid %s failed, err: %s", req.UUID, err)
			return nil, err
		}

		if user.UUID != uuid.Nil {
			if user.UUID != uuidUser.UUID {
				l.Errorf("look database have dirty data, %v %v", *user, *uuidUser)
				return nil, fmt.Errorf("internal error, database have dirty data")
			}
		}

		user = uuidUser
	}

	role, _ := user.Roles.String()
	var lockAt string
	if user.LockAt != nil {
		lockAt = user.LockAt.String()
	}

	resp = &types.GetUserResp{
		UUID:      user.UUID.String(),
		Username:  user.Username,
		NickName:  user.NickName,
		Roles:     role,
		LockAt:    lockAt,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return resp, nil
}
