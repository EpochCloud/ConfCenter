package config

import (
	"ConfCenter/log"
	"sync"
	"github.com/jmoiron/sqlx"
)

var(
	Log *log.Logger
	BufPool,Client *sync.Pool
	Db *sqlx.DB
	Conf *Config
)


type Config struct {
	LogConf LogConf
	HttpConf HttpConf
	MysqlConf MysqlConf
}

type LogConf struct {
	LogLevel string
	LogPath  string
}

type HttpConf struct {
	ReadTimeout int
	WriteTimeout  int
	IdleTimeout   int
	Addr       string
}

type MysqlConf struct {
	DriverName string
	Name string
	PassWord string
	Ip     string
	Port   string
	Database string
}


