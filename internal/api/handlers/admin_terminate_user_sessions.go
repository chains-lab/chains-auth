package handlers

import (
	"context"

	svc "github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/chains-lab/sso-svc/internal/api/responses"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) AdminTerminateUserSessions(ctx context.Context, req *svc.AdminTerminateUserSessionsRequest) (*svc.Empty, error) {
	requestID := uuid.New()
	meta := Meta(ctx)

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format user id: %s", req.UserId)
	}

	err = s.app.AdminTerminateSessions(ctx, meta.InitiatorID, userId)
	if err != nil {
		return nil, responses.AppError(ctx, requestID, err)
	}

	Log(ctx, requestID).Warnf("User sessions terminated by admin %s", meta.InitiatorID)

	return &svc.Empty{}, nil
}
