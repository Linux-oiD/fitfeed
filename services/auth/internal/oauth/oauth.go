package oauth

import (
	"fitfeed/auth/internal/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
)

func NewAuth(conf *config.AppConfig) {

	store := sessions.NewCookieStore([]byte(conf.Auth.Secret))
	store.MaxAge(conf.Auth.MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = conf.IsProd
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store

	providers := []goth.Provider{}

	for name, providerConf := range conf.Auth.Providers {
		if !providerConf.Enabled {
			continue // Skip disabled providers
		}

		var callbackURL string
		if conf.IsProd {
			callbackURL = fmt.Sprintf("%s://%s/%s/v1/oauth/%s/callback", conf.Web.Protocol, conf.Web.Hostname, conf.Auth.Prefix, name)
		} else {
			// For local development, use localhost with the port
			callbackURL = fmt.Sprintf("http://%s:%d/v1/oauth/%s/callback", conf.Web.Hostname, conf.Auth.Port, name)
		}

		log.Printf("Registering OAuth2 provider '%s' with callback URL: %s", name, callbackURL)

		switch name {
		case "google":
			providers = append(providers, google.New(providerConf.ClientID, providerConf.ClientSecret, callbackURL))
		case "yandex":
			providers = append(providers, yandex.New(providerConf.ClientID, providerConf.ClientSecret, callbackURL))
		case "github":
			// If you have other providers, add them here
			providers = append(providers, github.New(providerConf.ClientID, providerConf.ClientSecret, callbackURL))
		// Add more cases for other providers you support
		default:
			log.Printf("Warning: Provider '%s' is enabled but not recognized.", name)
		}
	}

	goth.UseProviders(providers...)
	log.Printf("Goth providers initialized. Total providers: %d", len(goth.GetProviders()))
}
