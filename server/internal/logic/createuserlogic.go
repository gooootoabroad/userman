package logic

import (
	"context"
	"fmt"

	"userman/server/global"
	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/model"

	"github.com/google/uuid"
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
	logx.Infof("req %v", req)
	username := req.Username
	// 判断是否已经存在username了
	userInfo := model.UserInfo{}
	global.DB.Where("Username = ?", username).First(&userInfo)
	if userInfo.Username != "" {
		return nil, fmt.Errorf("username %s exist", username)
	}

	if req.NickName != "" {
		// 有别名，检查别名
		global.DB.Where("NickName = ?", req.NickName).First(&userInfo)
		if userInfo.Username != "" {
			return nil, fmt.Errorf("nickname %s exist", req.NickName)
		}
	}

	// todo权限检查

	// 创建用户，记录到数据库中
	userInfo = model.UserInfo{
		UUID:     uuid.New(),
		Username: req.Username,
		NickName: req.NickName,
		Password: req.Password,
		Roles:    req.Roles,
	}

	global.DB.Create(userInfo)
	return &types.CreateUserResp{
		Message: "good",
	}, nil
}
