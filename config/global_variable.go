package config

import (
	"ConfCenter/basic"
	"ConfCenter/log"
	"sync"
)

var Log *log.Logger
var BufPool *sync.Pool

type Variable struct {
	LogVariable basic.Map
}
