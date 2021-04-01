package model

import "gorm.io/gorm"

const (
	RoleUser = "user"

	StatusActive = 1
	StatusInactive = 0
)

type CredUser struct {
	Role string
	HashedPassword string
	LoginName      string `gorm:"unique"`
	Status int
	gorm.Model
}
func (c *CredUser) GetCredUser() *CredUser {
	return c
}
