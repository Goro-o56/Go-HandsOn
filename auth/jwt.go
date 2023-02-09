package auth

import (
	"Go-Handson/entity"
	"context"
	_ "embed"
)

const (
	RoleKey     = "role"
	UserNameKey = "user_name"
)

type userIDKey struct {
}

func SetUserID(ctx context.Context, uid entity.UserID) context.Context {
	return context.WithValue(ctx, userIDKey{}, uid)
}

func GetUserID(ctx context.Context) (entity.UserID, bool) {
	id, ok := ctx.Value(userIDKey{}).(entity.UserID)
	return id, ok
}
