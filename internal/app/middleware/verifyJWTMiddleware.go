package middleware

import (
	"context"
	c "go-skeleton/internal/app/constant"
	"go-skeleton/internal/app/service"
	"go-skeleton/internal/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func VerifyJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := service.AuthService.ExtractJwtRequest(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			switch errors.Cause(err).(type) {
			case *util.UnauthorizedError, *jwt.ValidationError:
				util.ErrorResponse(w, http.StatusUnauthorized, err)
			default:
				util.ErrorResponse(w, http.StatusInternalServerError, err)
			}
		} else {

			ctx := context.WithValue(r.Context(), c.UserID, claims.ID)
			ctx = context.WithValue(ctx, c.UserRole, claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
