package LoginService

import (
	"kreditplus-api/jwt"
	"kreditplus-api/model/dto"
	"kreditplus-api/utils"
	"net/http"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
)

type loginService struct {
}

// LoginService implements LoginService.
func (e *loginService) LoginService(r *http.Request) (result dto.StandardResponse, err dto.StandardError) {
	
	var credentials jwt.Credentials

	err = utils.ReadBody(r, &credentials)

	if err.Err != nil {
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
        err = dto.StandardError{}.GenerateEmptyField("username or password")
        return
    }

	// real case harus nya user dari database
	if credentials.Username != "admin" || credentials.Password != "admin#1234" {
        err = dto.StandardError{}.GenerateInvalidCredentials()
        return
    }

	//set expired token misal 5 menit
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &jwt.Claims{
        Username: credentials.Username,
        StandardClaims: jwtLib.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
    tokenString, errS := token.SignedString([]byte(jwt.JWT_KEY))
    if errS != nil {
		err = dto.StandardError{
            Err:               errS,
            Code:              http.StatusInternalServerError,
            AdditionalMessage: "failed to generate jwt token",
        }
        return
    }


	result.Content = tokenString

	return
}

func NewLoginService() LoginService {
	return &loginService{}
}
