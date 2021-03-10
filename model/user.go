package model

import (
	"github.com/dealmaker/model/obj"
	"github.com/dealmaker/shared/base"
)

type UserInfoDomain struct {
	base.BaseDomain
	obj.UserInfo
}