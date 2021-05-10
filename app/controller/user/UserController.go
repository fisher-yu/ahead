package user

import (
	"ahead/app/controller"
	"ahead/app/service/user"
	"ahead/app/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.BaseController
	svc  user.UserService
}

// 用户列表
func (ctrl UserController) Index(ctx *gin.Context) {
	users, err := ctrl.svc.GetUserList()
	if err != nil {
		ctrl.Error(ctx, controller.ServerErrCode, "")
		return
	}

	ctrl.Success(ctx, users)
}

// 创建用户
func (ctrl UserController) Create(ctx *gin.Context) {
	params := validator.User{}
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctrl.Error(ctx, controller.ParamsErrCode, err.Error())
		return
	}
	userId, err := ctrl.svc.CreateUser(&params)
	if err != nil {
		ctrl.Error(ctx, controller.ServerErrCode, err.Error())
		return
	}
	result := make(map[string]int)
	result["user_id"] = userId

	ctrl.Success(ctx, result)
}

// 更新用户信息
func (ctrl UserController) Update(ctx *gin.Context) {
	params := validator.User{}
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctrl.Error(ctx, controller.ParamsErrCode, err.Error())
		return
	}
	if ok, err := ctrl.svc.UpdateUser(&params); !ok || err != nil {
		ctrl.Error(ctx, controller.ServerErrCode, err.Error())
		return
	}

	ctrl.Success(ctx, nil)
}

// 删除用户
func (ctrl UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctrl.Error(ctx, controller.ParamsErrCode, "")
		return
	}
	affectedRows, err := ctrl.svc.DeleteUser(userId)
	if affectedRows == 0 {
		ctrl.Error(ctx, controller.NotFoundErrCode, "")
		return
	}

	ctrl.Success(ctx, nil)
}

// 用户信息
func (ctrl UserController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctrl.Error(ctx, controller.ParamsErrCode, "")
		return
	}
	userInfo, err := ctrl.svc.GetUserInfo(userId)

	ctrl.Success(ctx, userInfo)
}
