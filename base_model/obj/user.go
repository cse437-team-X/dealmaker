package obj

type User struct {
	Username string
	Email string
	HashedPassword string
	Status int
}
func (u *User) GetHashedPassword() string {
	return u.HashedPassword
}
func (u *User) GetUsername() string {
	return u.Username
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetStatus() int {
	return u.Status
}