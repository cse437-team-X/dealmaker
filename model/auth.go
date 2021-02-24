package model

type UserLoginRequest struct {
	BaseRequest
	Username string
	HashedPassword string
}

func (u *UserLoginRequest) GetHashedPassword() string {
	return u.HashedPassword
}

func (u *UserLoginRequest) GetUsername() string {
	return u.Username
}


