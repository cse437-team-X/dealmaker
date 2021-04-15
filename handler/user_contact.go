package handler


import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	model3 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"gorm.io/gorm"
	"net/http"
)

type UserContactDomain struct {
	base.Base
	model.JwtAuth
	model2.ContactEmailInfo
	model2.EmailContent
}

type UserContactInput struct {
	Token string
	ItemTitle string
	UploaderId uint
}

func UserContact(c *gin.Context) {
	input := UserContactInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := UserContactDomain{
		JwtAuth: model.JwtAuth{
			Token: input.Token,
		},
		ContactEmailInfo: model2.ContactEmailInfo{
			To: model.CredUser{
				Model: gorm.Model{
					ID: input.UploaderId,
				},
			},
			Item: model3.Item{
				Title: input.ItemTitle,
			},
		},
	}

	s := factory.Factory.Get("/item/user/contact")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	c.JSON(code, nil)
}
