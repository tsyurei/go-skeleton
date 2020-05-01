package middleware

import (
	c "go-skeleton/internal/app/constant"
	eu "go-skeleton/internal/app/enum/user"

	"go-skeleton/internal/util"
	"net/http"
)

func VerifyAdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if userRole := r.Context().Value(c.UserRole); userRole != nil {
			if userRole != eu.Role.Admin {
				w.Header().Set("Content-Type", "application/json")
				util.ErrorResponse(w, http.StatusUnauthorized, util.NewUnauthorizedError().WithContext("Unauthorized"))
			} else {
				next.ServeHTTP(w, r)
			}

		} else {
			w.Header().Set("Content-Type", "application/json")
			util.ErrorResponse(w, http.StatusUnauthorized, util.NewUnauthorizedError().WithContext("Unauthorized"))
		}
	})
}
