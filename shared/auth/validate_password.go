package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ValidatePasswordInterface interface {
	GetCredUser() *model.CredUser
	GetBase() *base.Base
}

func (w *WorkerInstance) ValidatePassword(c *streamline.ConveyorBelt) int {
	inputCredUser := c.DataDomain.(ValidatePasswordInterface).GetCredUser()
	rbase := c.DataDomain.(ValidatePasswordInterface).GetBase()

	dbCredUser := w.FuncGetCredUser(inputCredUser)
	c.Debugw("cred_user form db", dbCredUser)

	if c.S.Action == "recover" {
		inputCredUser.Role = dbCredUser.Role
		inputCredUser.ID = dbCredUser.ID
		return http.StatusOK
	}
	if dbCredUser.HashedPassword != inputCredUser.HashedPassword {
		rbase.BaseMessage = "password and login name doesn't match"
		return http.StatusForbidden
	}

	if dbCredUser.Status != 1 {
		rbase.BaseMessage = "inactive user"
		return http.StatusForbidden
	}

	inputCredUser.Role = dbCredUser.Role
	inputCredUser.ID = dbCredUser.ID

	return http.StatusOK
}