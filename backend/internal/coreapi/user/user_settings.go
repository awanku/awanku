package user

import (
	"errors"
	"net/http"

	"github.com/awanku/awanku/internal/coreapi/contract"
	"github.com/awanku/awanku/internal/coreapi/utils/apihelper"
	"github.com/awanku/awanku/internal/coreapi/utils/ctxhelper"
	"github.com/awanku/awanku/pkg/core"
	"github.com/go-chi/render"
)

type UserSettingsService struct {
    UserSettingsStore contract.UserSettingsStore
}


func (s *UserSettingsService) Init() error {
    return nil
}

func (s *UserSettingsService) HandleGetSettings(w http.ResponseWriter, r *http.Request) {
    userID := ctxhelper.AuthenticatedUserID(r.Context())

    userSettings, err := s.UserSettingsStore.GetOrCreateByUserID(userID)
    if err != nil {
        apihelper.InternalServerErrResp(w, err)
        return
    }
    apihelper.JSON(w, http.StatusOK, userSettings.Settings)
}

func (s *UserSettingsService) HandlePatch(w http.ResponseWriter, r *http.Request) {
    userID := ctxhelper.AuthenticatedUserID(r.Context())

    userSettings, err := s.UserSettingsStore.GetOrCreateByUserID(userID)
    if err != nil {
        apihelper.InternalServerErrResp(w, err)
        return
    }

    settingsRequest := &UserSettingsRequest{}
    if err := render.Bind(r, settingsRequest); err != nil {
        apihelper.InternalServerErrResp(w, err)
        return
    }

    userSettings.Settings = settingsRequest.Settings

    settings, err := s.UserSettingsStore.Update(userSettings)
    if err != nil {
        apihelper.InternalServerErrResp(w, err)
        return
    }

    apihelper.JSON(w, http.StatusOK, settings)
}


type UserSettingsRequest struct {
    *core.UserSettings
}

func (u *UserSettingsRequest) Bind(r *http.Request) error {
    if u.Settings == nil {
        return errors.New("Missing required Settings fields.")
    }

    return nil
}
