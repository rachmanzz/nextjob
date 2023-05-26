package setup

import (
	"errors"
	"os"
	"strconv"
)

type VarArgType struct {
	BotAPI      *string
	ShowCaseAPI *string
	RedisHost   string
	RedisPass   string
	RedisDB     int
	Production  bool
}

var VarArgData *VarArgType

func PopulateVarArg() error {
	vararg := VarArgType{
		RedisHost:  "localhost:6379",
		RedisPass:  "",
		RedisDB:    0,
		Production: false,
	}

	botApi := os.Getenv("BOT_API")
	if botApi == "" {
		return errors.New("bot api is ampty")
	}

	vararg.BotAPI = &botApi

	showCaseApi := os.Getenv("SHOWCASE_API")

	if showCaseApi != "" {
		// empty is OK
		vararg.ShowCaseAPI = &showCaseApi
	}

	redisHost := os.Getenv("REDIS_HOST")

	// empty is OK
	if redisHost != "" {
		vararg.RedisHost = redisHost
	}

	redisPass := os.Getenv("REDIS_PASS")

	// empty is OK
	if redisPass != "" {
		vararg.RedisPass = redisPass
	}

	redisDB := os.Getenv("REDIS_DB")

	// empty is OK
	if redisDB != "" {
		db, err := strconv.Atoi(redisDB)
		if err == nil {
			vararg.RedisDB = db
		}

	}

	strProduction := os.Getenv("PRODUCTION")
	if strProduction != "" {
		isProd, err := strconv.ParseBool(strProduction)
		if err == nil && isProd {
			vararg.Production = isProd
		}

	}

	VarArgData = &vararg

	return nil
}
