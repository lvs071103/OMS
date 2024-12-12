package models

// ParamSignUp 注册请求的参数
type ParamSignUp struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Email           string `json:"email" binding:"required,email"`
	VerifyCode      string `json:"verify_code" binding:"required"`
}

// ParamLogin 登陆请求参数
type ParamsLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamList 获取列表请求参数
type ParamList struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 每页数据条数
}
