package logic

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	"userman/server/internal/svc"
	"userman/server/internal/types"
	"userman/server/model"
	"userman/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResq, err error) {
	// todo: add your logic here and delete this line
	username := req.Username
	// 查看用户是否存在
	userInfo, err := utils.GetUserByUsername(l.ctx, username)
	if err != nil {
		l.Errorf("get user %s info failed, err: %s", username, err)
		return nil, fmt.Errorf("user name or passwd is not correct")
	}

	// 校验密码
	dbpwd := userInfo.Password
	// 1. 分割密码字符串，找到加密随机数以及加密原文
	splitArr := strings.Split(dbpwd, model.Separator)
	if len(splitArr) != 3 {
		l.Errorf("user %s db pwd %s invaild", username, dbpwd)
		return nil, model.InternalError
	}

	srand := splitArr[1]
	encryptPwd := splitArr[2]
	// 2. 加随机数与用户传入的密码拼接进行sha256
	reqPwd := sha256.Sum256([]byte(srand + req.Password))
	reqSlice := reqPwd[:]
	reqString := fmt.Sprintf("%x", string(reqSlice))
	// 3. 对比sha值是否相等
	if reqString != encryptPwd {
		l.Errorf("user %s passwd %s not correct", username, req.Password)
		return nil, fmt.Errorf("user name or passwd is not correct")
	}

	// 生成token
	role, _ := userInfo.Roles.String()
	token, err := utils.CreateToken(l.ctx, username, userInfo.UUID, role)
	if err != nil {
		l.Errorf("create user %s token failed, err: %s", username, err)
		return nil, fmt.Errorf("internal error, please try again")
	}

	resp = &types.LoginResq{
		Token: *token,
	}

	return resp, nil
}
