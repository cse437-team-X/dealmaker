package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"gorm.io/gorm"
	"net/http"
)

type ActivateUserInterface interface {
	GetJwtAuth() *model.JwtAuth
}

func (w *WorkerInstance) ActivateUser(c *streamline.ConveyorBelt) int {
	input := c.DataDomain.(ActivateUserInterface).GetJwtAuth()
	uid := input.GetTokenClaim().Uid
	query := model.CredUser{
		Model: gorm.Model{ID: uid},
	}
	ret := w.FuncGetCredUser(&query)
	ret.Status = model.StatusActive
	err := w.FuncUpdateCredUser(ret)
	if err != nil {
		c.Errorw("Dal update err", err)
		return http.StatusInternalServerError
	}
	err = w.blockList.InvalidateToken(input.VToken.Token, input.VToken.StandardClaims)
	if err != nil {
		c.Errorw("invalidate token err", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
