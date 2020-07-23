package workspace

import (
	"net/http"

	"github.com/awanku/awanku/internal/coreapi/appctx"
	"github.com/awanku/awanku/internal/coreapi/utils/apihelper"
)

func HandleGetMyWorkspaces(w http.ResponseWriter, r *http.Request) {
	currentUser := appctx.AuthenticatedUser(r.Context())

	workspaces, err := getUserWorkspaces(r.Context(), currentUser.ID)
	if err != nil {
		apihelper.InternalServerErrResp(w, err)
		return
	}

	apihelper.JSON(w, http.StatusOK, workspaces)
}
