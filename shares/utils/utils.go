package utils

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const OrgKey contextKey = "org"

func SetOrg(ctx context.Context, org string) context.Context {
	return context.WithValue(ctx, OrgKey, org)
}

func GetOrg(ctx context.Context) (string, bool) {
	org, ok := ctx.Value(OrgKey).(string)
	return org, !ok
}

func GenerateUUID() string {
	return uuid.New().String()
}
