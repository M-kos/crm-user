package user

type UserStatus string

const (
	ACTIVE   UserStatus = "active"
	INACTIVE UserStatus = "inactive"
	DELETED  UserStatus = "deleted"
)

type UserFromRequest struct {
	Permissions int32
	UserType    string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Tg          string
	Password    string
	Status      UserStatus
}

type User struct {
	ID          string
	UUID        string
	Permissions int32
	UserType    string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Phone       string
	Tg          string
	Status      UserStatus
	CreatedAt   string
	UpdatedAt   string
}

func NewUser(newUser *UserFromRequest) *User {
	return &User{
		Permissions: newUser.Permissions,
		UserType:    newUser.UserType,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Phone:       newUser.Phone,
		Tg:          newUser.Tg,
		Status:      newUser.Status,
	}
}
