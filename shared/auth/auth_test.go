package auth

import (
	"fmt"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/mentally-friendly-infra/log"
	"github.com/itzmeerkat/streamline"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	d := struct {
		model.JwtAuth
		auth_db.UserCredModel
	}{}

	d.ID = 999
	d.UserCredModel.Role = "superadmin"

	c := streamline.ConveyorBelt{
		DataDomain: &d,
		S:          nil,
		Ctx:        nil,
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return "rua"
		},
	}
	SignTokenWithTokenExpireTime(time.Hour * 1000)(&c)
	fmt.Println(d.Token)
}

func TestValidate(t *testing.T) {
	tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjE5LCJSb2xlIjoidXNlciIsIlNjb3BlIjoibm9ybWFsIiwiaWF0IjoxNjE3MTUxMjk4LCJleHAiOjE2MTcxNTQ4OTh9.qmLhcBsWLqdONykHhDZfOsoToo_NFUkyYzw0DfVRbl8"
	d := model.JwtAuth{
		Token: tok,
	}

	c := streamline.ConveyorBelt{
		DataDomain: &d,
		S:          nil,
		Ctx:        nil,
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return "rua"
		},
	}
	r := Validate(&c)
	fmt.Println(r)
}