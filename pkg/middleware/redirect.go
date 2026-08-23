package middleware

import (
	"lookfar/internal/utils"
	"net/http"
)

func Redirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value == "" {
			utils.Logger().Error("Missing session cookie", "error", err)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
