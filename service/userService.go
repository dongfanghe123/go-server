package service

import (
	"context"
	"fmt"
	"go-server/dao"
	"go-server/model"
	"go-server/util"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UserRegisterByPhone(ctx context.Context, phone string) error {
	//1、查询数据库，判断是否存在同名用户

	if !util.IsValidPhone(phone) {
		fmt.Println("注册失败，手机号格式错误")
		return fmt.Errorf("注册失败，手机号格式错误")
	}

	params := make(map[string]interface{})
	params["phone"] = phone
	user, err := dao.DefaultUserDao.SelectUserByPhone(ctx, params)
	if err != nil {
		fmt.Println("SelectUserByPhone err:", err)
		return fmt.Errorf("SelectUserByPhone err: %w", err)
	}

	//用户已存在
	if user != nil {
		fmt.Printf("err: phone %s is exists\n", phone)
		return fmt.Errorf("err: phone %s is exists", phone)
	}

	// 2、创建用户对象

	now := time.Now()

	newUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("创建uuid错误")
		return fmt.Errorf("创建uuid错误")
	}

	s := strings.ReplaceAll(newUUID.String(), "-", "")
	nickName := "user_" + string([]rune(s)[0:11])

	newUser := &model.User{
		NickName:   &nickName,
		Phone:      &phone,
		CreateTime: &now,
		UpdateTime: &now,
	}

	// 3、保存用户

	rows, err := dao.DefaultUserDao.Save(ctx, newUser)
	if err != nil {
		fmt.Println("Save user err:", err)
		return fmt.Errorf("save user err: %w", err)
	}

	// 4、检查是否成功插入

	if rows != 1 {
		return fmt.Errorf("register user failed")
	}

	return nil

}
