package user

import (
	"net/http"

	"github.com/awanku/awanku/internal/coreapi/contract"
    "github.com/awanku/awanku/internal/coreapi/utils/ctxhelper"
    "github.com/awanku/awanku/internal/coreapi/utils/apihelper"
)

type UserSettingsService struct {
    UserSettingsStore contract.UserSettingsStore
}


func (s *UserSettingsService) Init() error {
    return nil
}

func (s *UserSettingsService) HandleGetSettings(w http.ResponseWriter, r *http.Request) {
    userID := ctxhelper.AuthenticatedUserID(r.Context())

    userSettings, err := s.UserSettingsStore.GetByUserID(userID)
    if err != nil {
        apihelper.InternalServerErrResp(w, err)
        return
    }
    apihelper.JSON(w, http.StatusOK, userSettings)
}
