package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type InsertUserInterface interface {
	GetCredUser() *model.CredUser
}

func (w *WorkerInstance) NewUser(c *streamline.ConveyorBelt) int {
	inputCredUser := c.DataDomain.(InsertUserInterface).GetCredUser()
	inputCredUser.Status = 0
	inputCredUser.Role = model.RoleUser

	err := w.FuncInsertCredUser(inputCredUser)
	if err != nil {
		c.Errorw("Dal error", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
