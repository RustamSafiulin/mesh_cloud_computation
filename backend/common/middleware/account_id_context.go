package middleware

import "context"

var contextKey = "account"

func NewAccountIDContext(ctx context.Context, accountID string) context.Context {
	return context.WithValue(ctx, contextKey, accountID)
}

func AccountIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKey).(string)
	return id, ok
}

func AccountIDMustFromContext(ctx context.Context) string {
	id, ok := ctx.Value(contextKey).(string)
	if !ok {
		panic("Account ID can't find in context.")
	}

	return id
}
