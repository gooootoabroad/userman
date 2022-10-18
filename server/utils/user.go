package utils

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"userman/server/global"
	"userman/server/model"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func CreateUsers(ctx context.Context, users []model.UserInfo) error {
	logger := logx.WithContext(ctx)
	errorResult := map[string]string{}
	var w sync.WaitGroup
	for i := 0; i < len(users); i++ {
		w.Add(1)
		go func(user model.UserInfo) {
			defer w.Done()
			err := CreateUser(ctx, user)
			if err != nil {
				errorResult[user.Username] = fmt.Sprintf("%s", err.Error())
			}
		}(users[i])
	}

	w.Wait()
	if len(errorResult) != 0 {
		errMsg := ""
		for k, v := range errorResult {
			errMsg += fmt.Sprintf("{%s: %s}", k, v)
		}
		return fmt.Errorf(errMsg)
	}

	logger.Infof("create users done")
	return nil
}

func CreateUser(ctx context.Context, user model.UserInfo) error {
	logger := logx.WithContext(ctx)
	// 检查用户名或者别名是否被注册了
	username := user.Username
	nickName := user.NickName
	oldUser := model.UserInfo{}
	err := global.DB.Where("Username = ?", username).First(&oldUser).Error
	if oldUser.Username != "" {
		logger.Errorf("create user failed, username %s  is already exist", oldUser.Username)
		return model.UserExist
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Errorf("find user %s info occured internal error %s", username, err)
		return model.InternalError
	}

	if nickName != "" {
		err = global.DB.Where("nick_name = ?", nickName).First(&oldUser).Error
		if oldUser.NickName != "" {
			logger.Errorf("create user failed, nick name %v is already exist", oldUser.NickName)
			return model.UserExist
		}

		if err != nil && err != gorm.ErrRecordNotFound {
			logger.Errorf("find user %s info occured internal error %s", username, err)
			return model.InternalError
		}
	}

	// 检查角色是否正确
	if _, err := user.Roles.String(); err != nil {
		logger.Errorf("user %s role %s is illegal, err:%s", username, user.Roles, err)
		return err
	}

	if user.UUID == uuid.Nil {
		user.UUID = uuid.New()
	}

	err = global.DB.Create(&user).Error
	if err != nil {
		logger.Errorf("create user %s failed from db, err %s", username, err)
		return err
	}

	logger.Infof("create user %s done", username)
	return nil
}

func GetUserByUsername(ctx context.Context, username string) (*model.UserInfo, error) {
	if username == "" {
		return nil, fmt.Errorf("need username")
	}

	logger := logx.WithContext(ctx)
	user := model.UserInfo{}
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Errorf("find user %s failed, err: %s", username, err)
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user %s is not exist", username)
		}

		return nil, err
	}

	return &user, nil
}

func GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*model.UserInfo, error) {
	if userUUID == uuid.Nil {
		return nil, fmt.Errorf("need uuid")
	}

	logger := logx.WithContext(ctx)
	user := model.UserInfo{}
	err := global.DB.Where("uuid = ?", userUUID).First(&user).Error
	if err != nil {
		logger.Errorf("find user uuid %s failed, err: %s", userUUID, err)
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user uuid %s is not exist", userUUID)
		}

		return nil, err
	}

	return &user, nil
}

func DeleteUser(ctx context.Context, username string, userUUID uuid.UUID) error {
	logger := logx.WithContext(ctx)
	// 检查用户是否存在，以及用户名和uuid映射关系
	user := model.UserInfo{}
	err := global.DB.Where("uuid = ?", userUUID).Find(&user).Error
	if err != nil {
		logger.Errorf("get user info by uuid %v failed, err: %s", userUUID, err)
		return err
	}

	if !strings.EqualFold(username, user.Username) {
		logger.Errorf("db username is %s requst username is %s", user.Username, username)
		return fmt.Errorf("username or uuid is not correct")
	}

	err = global.DB.Unscoped().Delete(&user).Error
	if err != nil {
		logger.Errorf("delete user %s failed, err: %s", username, err)
		return fmt.Errorf("internal error, please try again or check db")
	}

	return nil
}
