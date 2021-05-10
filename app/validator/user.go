package validator

type User struct {
	Id       int    `json:"id" form:"id"`
	Mobile   int64  `json:"mobile"  form:"mobile" binding:"required" `
	Email    string `json:"email"  form:"email"`
	Nickname string `json:"nickname"  form:"nickname" binding:"required"`
	Sex      int    `json:"sex"  form:"sex"`
	Status   int    `json:"status"  form:"status"`
	UserType int    `json:"user_type"  form:"user_type"`
}
