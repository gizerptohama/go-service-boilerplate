package constants

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_Issuer          string
	JWT_Secret          []byte
	JWT_Expire_Duration time.Duration
	JWT_Sign_Method     = jwt.SigningMethodHS256
)

func SetupJWTEnv() {
	JWT_Issuer = os.Getenv("APP_NAME")
	JWT_Secret = []byte(os.Getenv("SECRET_KEY"))
	JWTTimeOut, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		JWT_Expire_Duration = 30 * time.Minute
	}
	JWT_Expire_Duration = time.Duration(JWTTimeOut) * time.Minute
}
