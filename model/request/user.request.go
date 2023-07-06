package request

type UserCreateRequest struct {
	Username       string `json:"username" validate:"required"`
	Fullname       string `json:"fullname" validate:"required"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeatpassword" validate:"required"`
}
type UserUpdateRequest struct {
	Username       string `json:"username" `
	Password       string `json:"password" `
	RepeatPassword string `json:"repeatpassword" `
}