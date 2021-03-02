package model

type UserLoginDomain struct {
	BaseDomain
	Username string
	HashedPassword string
}
func (u *UserLoginDomain) GetHashedPassword() string {
	return u.HashedPassword
}
func (u *UserLoginDomain) GetUsername() string {
	return u.Username
}


type UserLogoutDomain struct {
	BaseDomain

}