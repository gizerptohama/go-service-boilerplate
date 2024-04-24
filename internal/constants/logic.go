package constants

import (
	"os"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

var (
	Zero_Decimal = decimal.NewFromInt(0)
	Code_Expire  time.Duration
)

func SetupLogicEnv() {
	exp, err := strconv.Atoi(os.Getenv("CODE_EXPIRE"))
	if err != nil {
		Code_Expire = 10 * time.Minute
	}
	Code_Expire = time.Duration(exp) * time.Minute
}
