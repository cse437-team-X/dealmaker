package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ValidatePasswordInterface interface {
	GetCredUser() *model.CredUser
}

func (w *WorkerInstance) ValidatePassword(c *streamline.ConveyorBelt) int {
	inputCredUser := c.DataDomain.(ValidatePasswordInterface).GetCredUser()
	dbCredUser := w.FuncGetCredUser(inputCredUser)
	if dbCredUser.Status != 1 ||
		dbCredUser.HashedPassword != inputCredUser.HashedPassword{
		return http.StatusForbidden
	}
	return http.StatusOK
}