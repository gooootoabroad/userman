package logic

import (
	"context"

	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// todo: add your logic here and delete this line
	uuid, err := uuid.Parse(req.UUID)
	if err != nil {
		l.Errorf("covenrsion uuid %s failed, err: %s", req.UUID, err)
		return nil, err
	}

	err = utils.DeleteUser(l.ctx, req.Username, uuid)
	if err != nil {
		l.Errorf("delete user %s failed, err: %s", req.Username, err)
		return nil, err
	}

	return &types.DeleteUserResp{
		Message: "success",
	}, nil
}
