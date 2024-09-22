package jwt

import "github.com/dgrijalva/jwt-go"

const JWT_KEY = "secret_key" // key sementara harusnya di letakan di env

//sample credentials
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}