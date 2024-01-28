// Code generated by goctl. DO NOT EDIT.
package types

type DetailRequest struct {
	UserID int64 `form:"user_id"`
}

type DetailResponse struct {
	Username string `json:"user_name"`
	Gender   int    `json:"gender"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	AccessExpire int    `json:"access_expire"`
	RefreshAfter int    `json:"refresh_after"`
}

type SignupRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required,min=4,max=32"`
	RePassword string `json:"re_password" validate:"required,min=4,max=32"`
	Gender     int    `json:"gender" validate:"oneof=0 1 2"`
}

type SignupResponse struct {
	Message string `json:"message"`
}
