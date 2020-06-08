package auth

import (
	"net/http"
	"strconv"

	"github.com/awanku/awanku/backend/internal/core/contracts"
	"github.com/awanku/awanku/backend/internal/core/utils/apihelper"
	"github.com/awanku/awanku/backend/pkg/model"
	"github.com/awanku/awanku/backend/pkg/oauth2provider"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Auth struct {
	Storage       contracts.UserStore
	config        *Config
	cookieManager *CookieManager

	providers map[string]contracts.AuthProvider
}

func (a *Auth) Init() error {
	a.providers = map[string]contracts.AuthProvider{
		"github": &oauth2provider.GithubProvider{
			Config: oauth2Config("development", "github"),
		},
		"google": &oauth2provider.GoogleProvider{
			Config: oauth2Config("development", "google"),
		},
	}
	a.cookieManager = newCookieManager("12345678901234561234567890123456", "1234567890123456")
	return nil
}

type getProviderConnectParam struct {
	Provider   string `json:"provider"`
	RedirectTo string `json:"redirect_to"`
	State      string `json:"state"`
	UserID     string `json:"user_id"`
}

func (p getProviderConnectParam) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Provider, validation.Required, validation.In("github", "google")),
		validation.Field(&p.RedirectTo, validation.Required, is.URL),
		validation.Field(&p.State, validation.Required),
		validation.Field(&p.UserID, is.Int),
	)
}

func (a *Auth) HandleGetProviderConnect(w http.ResponseWriter, r *http.Request) {
	param := &getProviderConnectParam{
		Provider:   chi.URLParam(r, "provider"),
		RedirectTo: r.URL.Query().Get("redirect_to"),
		State:      r.URL.Query().Get("state"),
		UserID:     r.URL.Query().Get("user_id"),
	}
	if err := param.Validate(); err != nil {
		apihelper.BadRequestErrResp(w, err)
		return
	}

	handler := a.providers[param.Provider]

	sessionData := map[string]string{
		"redirect_to": param.RedirectTo,
		"user_id":     param.UserID,
	}
	err := a.cookieManager.Put(w, "auth_data", sessionData)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}

	apihelper.RedirectResp(w, handler.LoginURL(param.State))
}

type getProviderCallbackParam struct {
	Provider string `json:"provider"`
	Code     string `json:"code"`
}

func (p getProviderCallbackParam) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Provider, validation.Required, validation.In("github", "google")),
		validation.Field(&p.Code, validation.Required),
	)
}

func (a *Auth) HandleGetProviderCallback(w http.ResponseWriter, r *http.Request) {
	param := &getProviderCallbackParam{
		Provider: chi.URLParam(r, "provider"),
		Code:     r.URL.Query().Get("code"),
	}
	if err := param.Validate(); err != nil {
		apihelper.BadRequestErrResp(w, err)
		return
	}

	handler := a.providers[param.Provider]

	userData, err := handler.ExchangeCode(param.Code)
	if err != nil {
		apihelper.BadRequestErrResp(w, map[string]string{
			"error": "failed to authenticate user",
		})
		return
	}

	sessionData, err := a.cookieManager.Fetch(r, "auth_data")
	if err != nil {
		apihelper.BadRequestErrResp(w, map[string]string{
			"error": "failed to fetch session data",
		})
		return
	}
	a.cookieManager.Destroy(w, "auth_data")

	targetUserID, _ := strconv.ParseInt(sessionData["user_id"], 10, 64)
	if targetUserID > 0 {
		user, err := a.Storage.FindByID(targetUserID)
		if err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
		if user.ID <= 0 {
			apihelper.BadRequestErrResp(w, map[string]map[string]interface{}{
				"errors": {
					"user_id": "does not exist",
				},
			})
			return
		}
		user.SetOauth2Identifier(userData.Provider, &userData.Identifier)
		if err := a.Storage.Save(user); err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
	} else {
		user := &model.User{
			Name:  userData.Name,
			Email: userData.Email,
		}
		user.SetOauth2Identifier(userData.Provider, &userData.Identifier)
		if err := a.Storage.Create(user); err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
	}

	apihelper.RedirectResp(w, sessionData["redirect_to"])
}

func (a *Auth) HandleGetToken(w http.ResponseWriter, r *http.Request) {

}
