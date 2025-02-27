package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/MudassirDev/all-in-one-business-tool/api/auth"
	"github.com/MudassirDev/all-in-one-business-tool/models"
)

func AuthMiddleware(next http.Handler, apiCfg *models.APICfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(apiCfg.AuthCookieName)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}

			log.Printf("Error retrieving cookie: %v", err)

			http.SetCookie(w, &http.Cookie{
				Name:   apiCfg.AuthCookieName,
				Value:  "",
				MaxAge: -1,
			})

			next.ServeHTTP(w, r)
			return
		}
		userId, err := auth.VerifyJWT(cookie.Value, apiCfg.JWTSecretKey)
		if err != nil {
			log.Println("jwt verifiation failed", err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), apiCfg.AuthCookieName, userId.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
