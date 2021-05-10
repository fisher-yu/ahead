package user

import (
	"ahead/app/model"
	"ahead/app/repository/user"
	"ahead/app/validator"
)

type UserService struct {
	repo *user.UserRepo
}

// 获取用户列表
func (svc *UserService) GetUserList() ([]model.User, error) {
	svc.repo = user.NewUserRepo()
	return svc.repo.GetUsers()
}

// 获取用户信息
func (svc *UserService) GetUserInfo(id int) (*model.User, error) {
	svc.repo = user.NewUserRepo()
	return svc.repo.GetUserById(id)
}

// 创建用户
func (svc *UserService) CreateUser(params *validator.User) (int, error) {
	svc.repo = user.NewUserRepo()
	return svc.repo.CreateUser(params)
}

// 更新用户信息
func (svc *UserService) UpdateUser(params *validator.User) (bool, error) {
	svc.repo = user.NewUserRepo()
	return svc.repo.UpdateUser(params)
}

// 更新用户信息
func (svc *UserService) DeleteUser(id int) (int64, error) {
	svc.repo = user.NewUserRepo()
	return svc.repo.DeleteUser(id)
}

