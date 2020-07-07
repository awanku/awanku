package user

import (
	"net/http"

	contracts "github.com/awanku/awanku/internal/coreapi/contract"
	"github.com/awanku/awanku/internal/coreapi/utils/apihelper"
	"github.com/awanku/awanku/internal/coreapi/utils/ctxhelper"
)

type UserService struct {
	UserStore contracts.UserStore
}

func (s *UserService) Init() error {
	return nil
}

// @Id api.v1.users.getMe
// @Summary Get current user data
// @Tags Users
// @Security ApiKeyAuth
// @Router /v1/users/me [get]
// @Produce json
// @Success 200 {object} core.User
// @Failure 400 {object} apihelper.HTTPError
// @Failure 401 {object} apihelper.HTTPError
// @Failure 500 {object} apihelper.InternalServerError
func (s *UserService) HandleGetMe(w http.ResponseWriter, r *http.Request) {
	userID := ctxhelper.AuthenticatedUserID(r.Context())
	user, err := s.UserStore.GetByID(userID)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}
	apihelper.JSON(w, http.StatusOK, user)
}
