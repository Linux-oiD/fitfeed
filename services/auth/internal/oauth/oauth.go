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

	if conf.Auth.Providers.google.Enabled {
		if conf.Auth.IsProd {
			callbackURL := fmt.Sprintf("%s://%s/%s/google/callback", conf.Web.Protocol, conf.Web.Hostname, conf.Auth.Prefix)
		} else {
			callbackURL := fmt.Sprintf("http://%s:%d/google/callback", conf.Web.Hostname, conf, Auth, Port)
		}
		goth.UseProviders(
			google.New(conf.Auth.Providers.google.ClientId, conf.Auth.Providers.google.ClientSecret, callbackURL),
		)
	}

}
