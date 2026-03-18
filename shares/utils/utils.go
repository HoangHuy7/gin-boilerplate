package utils

import "context"

type contextKey string

const OrgKey contextKey = "org"

func SetOrg(ctx context.Context, org string) context.Context {
	return context.WithValue(ctx, OrgKey, org)
}

func GetOrg(ctx context.Context) (string, bool) {
	org, ok := ctx.Value(OrgKey).(string)
	return org, !ok
}
