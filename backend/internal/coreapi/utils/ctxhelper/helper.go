package ctxhelper

import (
	"context"
)

type Key string

const (
	AuthenticatedUserIDKey Key = "authenticated_user_id"
)

func AuthenticatedUserID(ctx context.Context) int64 {
	user, _ := ctx.Value(AuthenticatedUserIDKey).(int64)
	return user
}
