package user

type User struct {
	ID           string `json:"id"`
	PermissionId string `json:"permission_id"`
	UserType     string `json:"user_type"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Tg           string `json:"tg"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func NewUser(id string, permissionId string, userType string, firstName string, lastName string, email string, phone string, tg string, createdAt string, updatedAt string) *User {
	return &User{
		ID:           id,
		PermissionId: permissionId,
		UserType:     userType,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Phone:        phone,
		Tg:           tg,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
