package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=3"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required"`
}
