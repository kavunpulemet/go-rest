package middlewares

import (
	"RESTAPIService2/pkg/api/utils"
	"RESTAPIService2/pkg/service/auth"
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "UserId"
)

type UserIdentityMiddleware struct {
	service auth.AuthorizationService
}

func NewUserIdentityMiddleware(service auth.AuthorizationService) *UserIdentityMiddleware {
	return &UserIdentityMiddleware{
		service: service,
	}
}

func (m *UserIdentityMiddleware) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := m.service.ParseToken(headerParts[1])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserCtx, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
