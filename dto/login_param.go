package dto

type LoginParam struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}
