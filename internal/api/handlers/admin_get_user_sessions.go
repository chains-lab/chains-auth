package handlers

import (
	"context"

	svc "github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/chains-lab/sso-svc/internal/api/responses"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) AdminGetUserSessions(ctx context.Context, req *svc.AdminGetUserSessionsRequest) (*svc.SessionsListResponse, error) {
	requestID := uuid.New()
	meta := Meta(ctx)

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID format: %v", err)
	}

	sessions, err := s.app.SelectUserSessions(ctx, userID)
	if err != nil {
		return nil, responses.AppError(ctx, requestID, err)
	}

	Log(ctx, requestID).Infof("retrieved sessions for user %s by admin %s", req.UserId, meta.InitiatorID)

	return responses.SessionList(sessions), nil
}
