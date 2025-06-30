package user

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	Permissions int32  `json:"permissions" validate:"required,numeric"`
	Type        string `json:"type" validate:"required,min=3"`
	FirstName   string `json:"firstName" validate:"required,min=2"`
	LastName    string `json:"lastName" validate:"required,min=2"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Phone       string `json:"phone,omitempty" validate:"e164,omitempty"`
	TgName      string `json:"tgName,omitempty" validate:"min=3,omitempty"`
	Status      string `json:"status" validate:"required,oneof=active inactive deleted"`
}

func (c *CreateUserRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

type GetUserByIdResponse struct {
	UUID        string `json:"uuid"`
	Permissions int32  `json:"permissions"`
	Type        string `json:"type"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone,omitempty"`
	TgName      string `json:"tgName,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateUserRequest struct {
	Permissions int32  `json:"permissions,omitempty" validate:"numeric,omitempty"`
	Type        string `json:"type,omitempty"  validate:"min=3,omitempty"`
	FirstName   string `json:"firstName,omitempty"  validate:"min=2,omitempty"`
	LastName    string `json:"lastName,omitempty"  validate:"min=2,omitempty"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password,omitempty" validate:"min=6,omitempty"`
	Phone       string `json:"phone,omitempty" validate:"e164,omitempty"`
	TgName      string `json:"tgName,omitempty" validate:"min=3,omitempty"`
	Status      string `json:"status,omitempty" validate:"oneof=active inactive deleted,omitempty"`
}

func (u *UpdateUserRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}
