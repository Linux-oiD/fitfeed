package oauth

import (
	"fitfeed/auth/internal/config"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewAuth(conf *config.AppConfig) {

	store := sessions.NewCookieStore([]byte(conf.Auth.Secret))
	store.MaxAge(conf.Auth.MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = conf.Auth.IsProd

	gothic.Store = store

	var callbackURL string

	for name, provider := range conf.Auth.Providers {

		if provider.Enabled {
			if conf.Auth.IsProd {
				callbackURL = fmt.Sprintf("%s://%s/%s/v1/oauth/%s/callback", conf.Web.Protocol, conf.Web.Hostname, conf.Auth.Prefix, name)
			} else {
				callbackURL = fmt.Sprintf("http://%s:%d/v1/oauth/%s/callback", conf.Web.Hostname, conf.Auth.Port, name)
			}
			goth.UseProviders(
				google.New(provider.ClientID, provider.ClientSecret, callbackURL),
			)
		}
	}
}
