package model

import (
	model2 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/auth/model"
)

type EmailContent struct {
	To        string
	Recipient string
	Title     string
	Body      string
}
func (e *EmailContent) GetEmailContent() *EmailContent {
	return e
}


type ContactEmailInfo struct {
	To model.CredUser
	Item model2.Item
}

func (c *ContactEmailInfo) GetContactEmailInfo() *ContactEmailInfo {
	return c
}