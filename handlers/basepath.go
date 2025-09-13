package handlers

import (
	"context"
	"net/http"
	"os"
	"strings"
)

type basePathContextKey struct{}

func withBasePath(basePath string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), basePathContextKey{}, basePath)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func basePathFromContext(ctx context.Context) string {
	val, ok := ctx.Value(basePathContextKey{}).(string)
	if !ok || val == "" {
		return "/"
	}
	return val
}

func normalizeBasePathFromEnv() string {
	raw := strings.TrimSpace(os.Getenv("PS_BASE_PATH"))
	if raw == "" || raw == "/" {
		return "/"
	}
	if !strings.HasPrefix(raw, "/") {
		raw = "/" + raw
	}
	// Remove trailing slash, except for root-only
	if strings.HasSuffix(raw, "/") {
		raw = strings.TrimRight(raw, "/")
		if raw == "" {
			return "/"
		}
	}
	return raw
}
