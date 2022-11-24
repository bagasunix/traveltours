package env

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Configs struct {
	Env              string
	JwtSecret        string
	JwtSecretRefresh string
	DbDriver         string
	DbHost           string
	DbPort           string
	DbName           string
	DbUsername       string
	DbPassword       string
	SSLMode          string
	Timezone         string
	MaxIdleConns     string
	MaxOpenConns     string
	ConnMaxLifetime  string
	LogPath          string
	LogName          string
	LogMaxSize       string
	LogMaxBackup     string
	LogMaxAge        string
}

var env = map[string]map[string]string{
	"ENV": {
		"Key":          "Env",
		"DefaultValue": "dev",
	},
	"JWT_SECRET_REFRESH": {
		"Key":          "JwtSecretRefresh",
		"DefaultValue": "GosNix",
	},
	"JWT_SECRET": {
		"Key":          "JwtSecret",
		"DefaultValue": "BagasUnix",
	},
	"DB_DRIVER": {
		"Key":          "DbDriver",
		"DefaultValue": "postgres",
	},
	"DB_HOST": {
		"Key":          "DbHost",
		"DefaultValue": "127.0.0.1",
	},
	"DB_PORT": {
		"Key":          "DbPort",
		"DefaultValue": "5432",
	},
	"DB_NAME": {
		"Key":          "DbName",
		"DefaultValue": "traveltour",
	},
	"DB_USER": {
		"Key":          "DbUsername",
		"DefaultValue": "bagasunix",
	},
	"DB_PASS": {
		"Key":          "DbPassword",
		"DefaultValue": "Kambing04",
	},
	"SSL_MODE": {
		"Key":          "SSLMode",
		"DefaultValue": "disable",
	},
	"TIME_ZONE": {
		"Key":          "Timezone",
		"DefaultValue": "Asia/Jakarta",
	},
	"MAX_IDLE_CONNS": {
		"Key":          "MaxIdleConns",
		"DefaultValue": "10",
	},
	"MAX_OPEN_CONNS": {
		"Key":          "MaxOpenConns",
		"DefaultValue": "100",
	},
	"CONN_MAX_LIFE_TIME": {
		"Key":          "ConnMaxLifetime",
		"DefaultValue": "5",
	},
	"LOG_PATH": {
		"Key":          "LogPath",
		"DefaultValue": "../../",
	},
	"LOG_NAME": {
		"Key":          "LogName",
		"DefaultValue": "boilerplate",
	},
	"LOG_MAXSIZE": {
		"Key":          "LogMaxSize",
		"DefaultValue": "10",
	},
	"LOG_MAXBACKUP": {
		"Key":          "LogMaxBackup",
		"DefaultValue": "5",
	},
	"LOG_MAXAGE": {
		"Key":          "LogMaxAge",
		"DefaultValue": "30",
	},
}

func LoadEnv() (*Configs, error) {
	logs, _ := zap.NewProduction()
	defer logs.Sync()

	var conf Configs

	err := godotenv.Load()

	if err != nil {
		logs.Fatal(err.Error())
	}

	for key, element := range env {
		if os.Getenv(key) != "" {
			reflect.ValueOf(&conf).Elem().FieldByName(element["Key"]).SetString(os.Getenv(key))
		} else {
			reflect.ValueOf(&conf).Elem().FieldByName(element["Key"]).SetString(element["DefaultValue"])
		}

	}
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
