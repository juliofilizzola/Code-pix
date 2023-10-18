package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Debug         = ""
	AutoMigrateDb = ""
	Env_Service   = ""
	Dns           = ""
	DbType        = ""
)

func LoadingEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Env_Service = fmt.Sprint(os.Getenv("ENVIRONMENT"))
	if Env_Service != "dev" {
		Dns = fmt.Sprint(os.Getenv("DNS"))
		DbType = os.Getenv("DB_TYPE")
	} else {
		Dns = fmt.Sprint(os.Getenv("DNS_DEV"))
		DbType = os.Getenv("DB_TYPE_DEV")
	}
	Port = fmt.Sprint(os.Getenv("PORT"))
	Debug = fmt.Sprint(os.Getenv("DEBUG"))
	AutoMigrateDb = fmt.Sprint(os.Getenv("AUTO_MIGRATE_DB"))
}
