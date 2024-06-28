package userdtos

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UpdateNameRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateEmailRequest struct {
	Email string `json:"email" form:"email" validate:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}
