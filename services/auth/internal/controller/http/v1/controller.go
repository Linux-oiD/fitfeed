package v1

import (
	"fitfeed/auth/internal/usecase"
)

type V1 struct {
	u usecase.UserManager
	o usecase.OauthManager
	p usecase.ProfileManager
}
