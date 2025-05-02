package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/response"
	"os"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

const (
	CookieName       = "oneiot_token"
	ClaimsContextKey = contextKey("userClaims")
)

// JWTMiddleware verifies the JWT token from the cookie
func JWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie(CookieName)
		if err != nil {

			if err == http.ErrNoCookie {
				UnauthorizedResponse(w, "Authorization cookie not found")
				return
			}
			http.Error(w, "Internal server error reading cookie", http.StatusInternalServerError)
			return
		}

		tokenString := cookie.Value
		if strings.TrimSpace(tokenString) == "" {
			UnauthorizedResponse(w, "Authorization token is empty")
			return
		}

		claims, err := helper.ValidateJWT(tokenString)
		if err != nil {
			ClearAuthCookie(w)
			UnauthorizedResponse(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		r = r.WithContext(ctx)

		next(w, r, ps)
	}
}

// GetClaimsFromContext retrieves the validated claims from the request context
func GetClaimsFromContext(ctx context.Context) (*helper.AppClaims, bool) {
	claims, ok := ctx.Value(ClaimsContextKey).(*helper.AppClaims)
	return claims, ok
}

// unauthorizedResponse sends a standard 401 response
func UnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response.SimpleResponse{
		Message: "Unauthorized: " + message,
		Data:    nil,
	})
}

// clearAuthCookie sends instructions to the browser to delete the auth cookie
func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expire immediately
		MaxAge:   -1,              // Tell browser to delete it
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") == "production", // Use Secure in production
		SameSite: http.SameSiteLaxMode,                 // Or StrictMode
	})
}
