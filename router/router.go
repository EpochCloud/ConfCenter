package router

import (
	"net/http"
	"ConfCenter/engine"
)

func Router(mux *http.ServeMux){
	mux.HandleFunc("/service_operation",engine.OperationService)
	mux.HandleFunc("gateway_configuration",engine.Configuration)
	mux.HandleFunc("/",engine.Hello)
}
