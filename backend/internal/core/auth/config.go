package auth

import (
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
	oauth2google "golang.org/x/oauth2/google"
)

type Config struct {
	BaseURL string
}

func oauth2Config(environment, provider string) *oauth2.Config {
	data := map[string]map[string]*oauth2.Config{
		"development": {
			"github": &oauth2.Config{
				ClientID:     "0c30bd255139cd20d33d",
				ClientSecret: "24c537a65760293ef284d84a3b1938dd16097825",
				Scopes:       []string{"read:user", "user:email"},
				Endpoint:     oauth2github.Endpoint,
				RedirectURL:  "http://api.awanku.xyz/v1/auth/github/callback",
			},
			"google": &oauth2.Config{
				ClientID:     "757848106543-b069r475lcql7373vmhk3179u5l1anek.apps.googleusercontent.com",
				ClientSecret: "R_JRM20ol-YFqzbVilo81sey",
				Scopes: []string{
					"https://www.googleapis.com/auth/userinfo.email",
					"https://www.googleapis.com/auth/userinfo.profile",
				},
				Endpoint:    oauth2google.Endpoint,
				RedirectURL: "http://api.awanku.xyz/v1/auth/google/callback",
			},
		},
	}
	return data[environment][provider]
}
