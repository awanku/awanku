package appctx

import (
	"context"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/pkg/core"
)

// Key context key
type Key string

// context keys
const (
	KeyEnvironment       Key = "environment"
	KeyDatabase          Key = "database"
	KeyAuthenticatedUser Key = "authenticated_user"
)

// Environment fetch environment name from context
func Environment(ctx context.Context) string {
	raw := ctx.Value(KeyEnvironment)
	if val, ok := raw.(string); ok {
		return val
	}
	return ""
}

// Database fetch database instance from context
func Database(ctx context.Context) *hansip.Cluster {
	raw := ctx.Value(KeyDatabase)
	if val, ok := raw.(*hansip.Cluster); ok {
		return val
	}
	return nil
}

// AuthenticatedUser fetch authenticated user from context
func AuthenticatedUser(ctx context.Context) *core.User {
	raw := ctx.Value(KeyAuthenticatedUser)
	if val, ok := raw.(*core.User); ok {
		return val
	}
	return nil
}
