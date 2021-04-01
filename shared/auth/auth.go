package auth

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"github.com/kataras/jwt"
	"net/http"
	"time"
)

type WorkerInstance struct {
	FuncGetCredUser func(*model.CredUser) *model.CredUser
	FuncInsertCredUser func(*model.CredUser) error
	FuncUpdateCredUser func(*model.CredUser) error

	InvalidTokenForgetTime time.Duration
	TokenExpireTimes map[string]time.Duration

	sharedKey []byte
	blockList *jwt.Blocklist
}

func (w WorkerInstance) Init() WorkerInstance {
	w.blockList = jwt.NewBlocklist(w.InvalidTokenForgetTime)
	w.sharedKey = []byte("p@ssw0rd")
	if _,ok := w.TokenExpireTimes[model.JwtScopeNormal]; !ok {
		w.TokenExpireTimes[model.JwtScopeNormal] = time.Hour
	}
	if _,ok := w.TokenExpireTimes[model.JwtScopeActivate]; !ok {
		w.TokenExpireTimes[model.JwtScopeActivate] = time.Hour * 24
	}
	if _,ok := w.TokenExpireTimes[model.JwtScopeRecover]; !ok {
		w.TokenExpireTimes[model.JwtScopeRecover] = time.Hour * 24
	}
	return w
}

func (w *WorkerInstance) Logout(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(model.JwtInterface).GetJwtAuth()
	token := data.VToken
	err := w.blockList.InvalidateToken(token.Token, token.StandardClaims)
	if err != nil {
		panic(err)
	}
	return http.StatusOK
}
//
//func (w *WorkerInstance) SetRecover(c *streamline.ConveyorBelt) int {
//	data := c.DataDomain.(auth_db.AuthDBInterface).GetUserCredModel()
//	data.Role = "recover"
//	return http.StatusOK
//}
