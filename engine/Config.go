package engine

import (
	"ConfCenter/config"
	"ConfCenter/model/mysql"
)

//-----------------------model
var (
	getwayManager = mysql.NewGatewayManager()
	service       = mysql.NewService()
	allService    = mysql.NewAllService()
)

//----------------------HTTP
var (
	result       = config.NewResult()
	normalResult = &config.NormalResult{}
	errResult    = config.NewErrorResult()
)

//--------------------
const scheme = "http://"
