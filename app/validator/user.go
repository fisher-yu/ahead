package validator

type UserCreate struct {
	Mobile   int64  `json:"mobile"  form:"mobile" binding:"required" `
	Nickname string `json:"nickname"  form:"nickname" binding:"required"`
	Email    string `json:"email"  form:"email"`
	Sex      int    `json:"sex"  form:"sex"`
	Status   int    `json:"status"  form:"status"`
	UserType int    `json:"user_type"  form:"user_type"`
}

type UserUpdate struct {
	Id       int    `json:"id" form:"id"  binding:"required" `
	Mobile   int64  `json:"mobile"  form:"mobile"`
	Nickname string `json:"nickname"  form:"nickname"`
	Email    string `json:"email"  form:"email"`
	Sex      int    `json:"sex"  form:"sex"`
	Status   int    `json:"status"  form:"status"`
	UserType int    `json:"user_type"  form:"user_type"`
}

type UserDelete struct {
	Id       int    `json:"id" form:"id"  binding:"required" `
}
