package dao

import (
	"context"
	"database/sql"
	"errors"
	"go-server/model"

	"github.com/go-juicedev/juice"
)

type UserDao interface {
	SelectByName(ctx context.Context, params map[string]interface{}) (*model.User, error)
	Save(ctx context.Context, user *model.User) (sql.Result, error)
	SelectUserByPhone(ctx context.Context, params map[string]interface{}) (*model.User, error)
}

type UserDaoImpl struct {
	manager juice.Manager
}

func (u *UserDaoImpl) SelectByName(ctx context.Context, params map[string]interface{}) (*model.User, error) {
	executor := juice.NewGenericManager[*model.User](u.manager).Object(UserDao(u).SelectByName)
	user, err := executor.QueryContext(ctx, params)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (u *UserDaoImpl) SelectUserByPhone(ctx context.Context, params map[string]interface{}) (*model.User, error) {
	executor := juice.NewGenericManager[*model.User](u.manager).Object(UserDao(u).SelectUserByPhone)
	user, err := executor.QueryContext(ctx, params)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

//func (u *UserDaoImpl) Save(ctx context.Context, params map[string]interface{}) (int64, error) {
//	executor := juice.NewGenericManager[*model.User](u.manager).Object(UserDao(u).Save)
//	result, err := executor.ExecContext(ctx, params)
//	if err != nil {
//		return -1,err
//	}
//
//	rows, err := result.RowsAffected()
//	if err != nil {
//		fmt.Println("RowsAffected error: err",err)
//		return -1,err
//	}
//
//	fmt.Println("影响行数：", rows)
//	return rows,nil
//}

func (u *UserDaoImpl) Save(ctx context.Context, user *model.User) (sql.Result, error) {
	executor := juice.NewGenericManager[*model.User](u.manager).
		Object(UserDao(u).Save)

	result, err := executor.ExecContext(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserDao() UserDao {
	return &UserDaoImpl{
		manager: GetEngine(), // 从全局获取
	}
}
