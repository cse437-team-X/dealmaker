package obj

type UserInfo struct {
	Username string        `gorm:"unique"`
	Email string           `gorm:"unique"`
	HashedPassword string
	Status int
}
func (u *UserInfo) GetHashedPassword() string {
	return u.HashedPassword
}
func (u *UserInfo) GetUsername() string {
	return u.Username
}
func (u *UserInfo) GetEmail() string {
	return u.Email
}
func (u *UserInfo) GetStatus() int {
	return u.Status
}