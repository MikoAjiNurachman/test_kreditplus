package config

import (
	"database/sql"
	"os"
)

type serverConfig struct {
	DBCon      *sql.DB
	ServerPort string
}

var ServerConfig = serverConfig{}

func InitConfiguration() {
	ServerConfig.DBCon = initDB()
	err := initMigration(ServerConfig.DBCon)
	if err != nil {
		return
	}
	ServerConfig.ServerPort = os.Getenv("SERVER_PORT")
}
