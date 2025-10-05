package middleware

import (
	"context"
	"go-api-native/helper"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			helper.Response(w, 401, "unauthorized", nil)
			return
		}

		user, err := helper.ValidateToken(accessToken)

		if err != nil {
			helper.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
