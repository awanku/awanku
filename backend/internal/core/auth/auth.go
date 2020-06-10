package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/schema"

	"github.com/awanku/awanku/backend/internal/core/contracts"
	"github.com/awanku/awanku/backend/internal/core/utils"
	"github.com/awanku/awanku/backend/internal/core/utils/apihelper"
	"github.com/awanku/awanku/backend/pkg/model"
	"github.com/awanku/awanku/backend/pkg/oauth2provider"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const authorizationCodeBufferLength = 15
const oauthTokensLength = 20

type AuthService struct {
	UserStore           contracts.UserStore
	AuthStore           contracts.AuthStore
	OauthTokenSecretKey []byte

	config        *Config
	cookieManager *CookieManager

	providers map[string]contracts.AuthProvider
}

func (a *AuthService) Init() error {
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

	userStore contracts.UserStore
}

func (p getProviderConnectParam) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Provider, validation.Required, validation.In("github", "google")),
		validation.Field(&p.RedirectTo, validation.Required, is.URL),
		validation.Field(&p.State, validation.Required),
		validation.Field(&p.UserID, is.Int, validation.By(p.validateUserID())),
	)
}

func (p getProviderConnectParam) validateUserID() validation.RuleFunc {
	return func(value interface{}) error {
		valueStr, _ := value.(string)
		if valueStr == "" {
			return nil
		}

		userID, _ := strconv.ParseInt(valueStr, 10, 64)
		user, err := p.userStore.FindByID(userID)
		if err != nil {
			return validation.NewInternalError(err)
		}
		if user.ID > 0 {
			return nil
		}
		return errors.New("does not exist")
	}
}

func (a *AuthService) HandleGetProviderConnect(w http.ResponseWriter, r *http.Request) {
	param := &getProviderConnectParam{
		Provider:   chi.URLParam(r, "provider"),
		RedirectTo: r.URL.Query().Get("redirect_to"),
		State:      r.URL.Query().Get("state"),
		UserID:     r.URL.Query().Get("user_id"),
		userStore:  a.UserStore,
	}
	if err := param.Validate(); err != nil {
		apihelper.ValidationErrResp(w, err)
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

func (a *AuthService) HandleGetProviderCallback(w http.ResponseWriter, r *http.Request) {
	param := &getProviderCallbackParam{
		Provider: chi.URLParam(r, "provider"),
		Code:     r.URL.Query().Get("code"),
	}
	if err := param.Validate(); err != nil {
		apihelper.ValidationErrResp(w, err)
		return
	}

	handler := a.providers[param.Provider]

	userData, err := handler.ExchangeCode(param.Code)
	if err != nil {
		apihelper.ValidationErrResp(w, map[string]string{
			"code": "invalid",
		})
		return
	}

	sessionData, err := a.cookieManager.Fetch(r, "auth_data")
	if err != nil {
		apihelper.BadRequestErrResp(w, "invalid_cookie", map[string]string{
			"cookie": "malformed format",
		})
		return
	}
	a.cookieManager.Destroy(w, "auth_data")

	var user *model.User
	targetUserID, _ := strconv.ParseInt(sessionData["user_id"], 10, 64)
	if targetUserID > 0 {
		user, err = a.UserStore.FindByID(targetUserID)
		if err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
		if user.ID <= 0 {
			apihelper.ValidationErrResp(w, map[string]string{
				"user_id": "does not exist",
			})
			return
		}
		user.SetOauth2Identifier(userData.Provider, &userData.Identifier)
		if err := a.UserStore.Save(user); err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
	} else {
		user = &model.User{
			Name:  userData.Name,
			Email: userData.Email,
		}
		user.SetOauth2Identifier(userData.Provider, &userData.Identifier)
		if err := a.UserStore.FindOrCreateByEmail(user); err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
	}

	codeStr, err := buildAuthorizationCode(authorizationCodeBufferLength)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}

	code, err := a.AuthStore.CreateAuthorizationCode(user.ID, codeStr)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}

	redirectURL, _ := url.Parse(sessionData["redirect_to"])
	query := redirectURL.Query()
	query.Add("code", code.Code)
	redirectURL.RawQuery = query.Encode()
	apihelper.RedirectResp(w, redirectURL.String())
}

type postTokenParam struct {
	GrantType    string `json:"grant_type" schema:"grant_type"`
	Code         string `json:"code" schema:"code"`
	RefreshToken string `json:"refresh_token" schema:"refresh_token"`

	retrievedCode       *model.OauthAuthorizationCode
	retrievedOauthToken *model.OauthToken

	authStore           contracts.AuthStore
	oauthTokenSecretKey []byte
}

func (p *postTokenParam) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.GrantType, validation.Required, validation.In("authorization_code", "refresh_token")),
		validation.Field(&p.Code, validation.By(p.validateCode())),
		validation.Field(&p.RefreshToken, validation.By(p.validateRefreshToken())),
	)
}

func (p *postTokenParam) validateCode() validation.RuleFunc {
	return func(value interface{}) error {
		if p.GrantType != "authorization_code" {
			return nil
		}

		code, _ := value.(string)
		if p.GrantType == "authorization_code" && value == "" {
			return validation.ErrRequired
		}

		var err error
		p.retrievedCode, err = p.authStore.GetAuthorizationCodeByCode(code)
		if err != nil {
			return validation.NewInternalError(err)
		}

		if p.retrievedCode.Code != "" {
			return nil
		}
		return errors.New("invalid")
	}
}

func (p *postTokenParam) validateRefreshToken() validation.RuleFunc {
	return func(value interface{}) error {
		if p.GrantType != "refresh_token" {
			return nil
		}

		tokenRaw, _ := value.(string)
		if p.GrantType == "refresh_token" && value == "" {
			return validation.ErrRequired
		}

		tokenParts := strings.Split(tokenRaw, ":")
		if len(tokenParts) != 2 {
			return errors.New("invalid")
		}

		tokenIDStr := tokenParts[0]
		refreshTokenDecoded, err := base64.URLEncoding.DecodeString(tokenParts[1])
		if err != nil {
			return errors.New("invalid")
		}

		tokenID, _ := strconv.ParseInt(tokenIDStr, 10, 64)
		p.retrievedOauthToken, err = p.authStore.GetOauthToken(tokenID)
		if err != nil {
			return validation.NewInternalError(err)
		}

		valid, err := utils.ValidateHmac(p.oauthTokenSecretKey, refreshTokenDecoded, p.retrievedOauthToken.RefreshTokenHash)
		if err != nil {
			return validation.NewInternalError(err)
		}
		if valid {
			return nil
		}
		return errors.New("invalid")
	}
}

func (a *AuthService) HandlePostToken(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
			"request_body": "malformed format",
		})
		return
	}

	param := postTokenParam{
		authStore:           a.AuthStore,
		oauthTokenSecretKey: a.OauthTokenSecretKey,
	}
	err := schema.NewDecoder().Decode(&param, r.PostForm)
	if err != nil {
		apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
			"request_body": "malformed format",
		})
		return
	}
	if err := param.Validate(); err != nil {
		apihelper.ValidationErrResp(w, err)
		return
	}

	token, err := buildOauthToken(a.OauthTokenSecretKey, oauthTokensLength)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}
	token.RequesterIP = r.Header.Get("X-Forwarded-For")
	token.RequesterUserAgent = r.Header.Get("User-Agent")

	switch param.GrantType {
	case "refresh_token":
		// if grant type is refresh_token, also delete old token
		if err := a.AuthStore.DeleteOauthToken(param.retrievedOauthToken.ID); err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}
		token.UserID = param.retrievedOauthToken.UserID
	case "authorization_code":
		token.UserID = param.retrievedCode.UserID
	}

	if err := a.AuthStore.CreateOauthToken(token); err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}
	apihelper.JSON(w, http.StatusOK, token.Token())
}