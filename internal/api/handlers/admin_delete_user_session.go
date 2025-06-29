package handlers

import (
	"context"

	svc "github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/chains-lab/sso-svc/internal/api/responses"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) AdminDeleteUserSession(ctx context.Context, req *svc.AdminDeleteUserSessionRequest) (*svc.Empty, error) {
	requestID := uuid.New()
	meta := Meta(ctx)

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		Log(ctx, requestID).WithError(err).Errorf("invalid format user id: %s", req.UserId)

		return nil, status.Errorf(codes.InvalidArgument, "invalid format user id: %s", req.UserId)
	}

	err = s.app.DeleteSession(ctx, meta.InitiatorID, userId)
	if err != nil {
		Log(ctx, requestID).WithError(err).Error("failed to delete session")

		return nil, responses.AppError(ctx, requestID, err)
	}

	Log(ctx, requestID).WithField("user_id", userId).Infof("User sessions deleted by admin %s", meta.InitiatorID)
	return &svc.Empty{}, nil
}
