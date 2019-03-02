package router

import (
	"net/http"
	"ConfCenter/engine"
)

/*
/apigateway_operation
{
	"route": "/",
	"service": {
		"serviceaddr": ["192.168.56.11:8080", "192.168.56.11:8081"],
		"registertime": "2019.2.22",
		"altreason": "xx"
	},
	"servicename": "aa"
}

/gateway_configuration
{
	"ip":"127.0.0.1",
	"port":"8090",
	"timeout":15,
	"loglevel":"debug",
	"logpath":"D:/project/src/quick/logcatlog",
	"bufpool":0,
	"intranetip":"127.0.0.1",
	"intranetport":"6060",
	"maxheader":"",
	"managerroute":"/manager",
	"serviceroute":"/service"
}
*/

func Router(mux *http.ServeMux){
	mux.HandleFunc("/quick_operation",engine.OperationService)
	mux.HandleFunc("/quick_configuration",engine.Configuration)
	mux.HandleFunc("/service_operation",engine.Operation)
	mux.HandleFunc("/pull_operation",engine.PullOperation)
}
