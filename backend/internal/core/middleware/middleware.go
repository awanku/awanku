package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/awanku/awanku/backend/internal/core/contracts"
	"github.com/awanku/awanku/backend/internal/core/utils"
	"github.com/awanku/awanku/backend/internal/core/utils/apihelper"
	"github.com/awanku/awanku/backend/internal/core/utils/ctxhelper"
)

type Middleware struct {
	OauthTokenSecretKey []byte
	AuthStore           contracts.AuthStore
}

func (m *Middleware) ValidateOauthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerParts := strings.Split(r.Header.Get("authorization"), " ")
		if len(headerParts) != 2 {
			apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
				"authorization_header": "malformed format",
			})
			return
		}

		tokenType := headerParts[0]
		if strings.ToLower(tokenType) != "bearer" {
			apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
				"authentication_type": "not supported",
			})
			return
		}

		accessTokenRaw := headerParts[1]
		accessTokenParts := strings.Split(accessTokenRaw, ":")
		if len(accessTokenParts) != 2 {
			apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
				"authentication_credentials": "malformed format",
			})
			return
		}

		accessTokenID := accessTokenParts[0]
		decodedAccessToken, err := base64.URLEncoding.DecodeString(accessTokenParts[1])
		if err != nil {
			apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
				"authentication_credentials": "malformed format",
			})
			return
		}

		tokenIDInt, _ := strconv.ParseInt(accessTokenID, 10, 64)
		if tokenIDInt <= 0 {
			apihelper.BadRequestErrResp(w, "invalid_request", map[string]string{
				"authentication_credentials": "malformed format",
			})
			return
		}

		storedToken, err := m.AuthStore.GetOauthToken(tokenIDInt)
		if err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}

		valid, err := utils.ValidateHmac(m.OauthTokenSecretKey, []byte(decodedAccessToken), storedToken.AccessTokenHash)
		if err != nil {
			apihelper.InternalServerErrResp(w, err)
			return
		}

		if !valid {
			apihelper.UnauthorizedAccessResp(w, "access_denied", map[string]string{
				"access_token": "invalid",
			})
			return
		}

		ctx := context.WithValue(r.Context(), ctxhelper.AuthenticatedUserIDKey, storedToken.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
