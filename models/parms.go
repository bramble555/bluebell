package models

type ParamSignUp struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	// binding里面要求的字段名不是json字段名，是结构体里面的字段名
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
