package MiddlewareHandler

import (
	"kreditplus-api/jwt"
	jwtLib "github.com/dgrijalva/jwt-go"
	"net/http"
)

func TokenVerifyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Missing authorization header", http.StatusUnauthorized)
            return
        }

        claims := &jwt.Claims{} // bisa disesuaikan kondisi claims jwt
        token, err := jwtLib.ParseWithClaims(tokenString, claims, func(token *jwtLib.Token) (interface{}, error) {
            return []byte(jwt.JWT_KEY), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}