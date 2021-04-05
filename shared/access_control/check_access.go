package access_control

import (
	"fmt"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type CheckAccessInterface interface {
	GetJwtAuth() *model.JwtAuth
	GetBase() *base.Base
}

func (w *WorkerInstance) CheckAccess(c *streamline.ConveyorBelt) int {
	a := c.DataDomain.(CheckAccessInterface).GetJwtAuth()
	b := c.DataDomain.(CheckAccessInterface).GetBase()
	rbacSubject := fmt.Sprintf("%v_%v", a.Scope, a.Role)
	rbacObject := c.S.Resource
	rbacAction := c.S.Action

	ok, err := w.enforcer.Enforce(rbacSubject, rbacObject, rbacAction)
	if err != nil {
		b.BaseMessage = err.Error()
		return http.StatusInternalServerError
	}
	if !ok {
		b.BaseMessage = fmt.Sprintf("%v is not allowed for %v %v", rbacSubject, rbacAction, rbacObject)
		return http.StatusForbidden
	}
	return http.StatusOK
}
