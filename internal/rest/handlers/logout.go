package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/google/uuid"
)

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	user, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		h.presenters.InvalidToken(w, requestID, err)
		return
	}

	appErr := h.app.Logout(r.Context(), user.SessionID)
	if appErr != nil {
		h.presenters.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Infof("logout session id: %s", user.SessionID)
	httpkit.Render(w, http.StatusNoContent)
}
