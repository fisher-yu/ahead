package user

import (
	"ahead/app/model"
	"ahead/app/repository"
	"ahead/app/validator"
	"ahead/core"
)

type UserRepo struct {
	*core.Model
}

func NewUserRepo() *UserRepo {
	return &UserRepo{core.NewModel()}
}

// 获取用户列表
func (repo *UserRepo) GetUsers() ([]model.User, error) {
	var users = make([]model.User, 0)
	err := repo.DB.Omit("auth_key", "password_hash").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 获取用户信息
func (repo *UserRepo) GetUserById(id int) (*model.User, error) {
	var user = model.User{}
	ok, err := repo.DB.Omit("auth_key", "password_hash").ID(id).Get(&user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &user, nil
}

// 创建用户
func (repo *UserRepo) CreateUser(vUser *validator.User) (int, error) {
	mUser := model.User{}
	base := &repository.BaseRepo{}
	base.BindModel(vUser, &mUser)
	_, err := repo.DB.Insert(&mUser)
	if err != nil {
		return 0, err
	}

	return mUser.Id, nil
}

// 更新用户
func (repo *UserRepo) UpdateUser(vUser *validator.User) (bool, error) {
	mUser := model.User{}
	base := &repository.BaseRepo{}
	base.BindModel(vUser, &mUser)
	_, err := repo.DB.ID(mUser.Id).Update(&mUser)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 更新用户
func (repo *UserRepo) DeleteUser(id int) (int64, error) {
	mUser := model.User{}
	return repo.DB.ID(id).Delete(&mUser)
}
