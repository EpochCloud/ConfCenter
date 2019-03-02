package engine

import (
	"net/http"
	"ConfCenter/config"
	"ConfCenter/basic"
	"encoding/json"
	"fmt"
	"errors"
	"time"
)

func PullOperation(w http.ResponseWriter,r *http.Request){
	switch  {
	case r.Method == http.MethodGet:
		err := pushOperation(w,r)
		if err != nil {
			return
		}
		return
	default:
		errResult.SendErrorResponse(w, config.ErrorMethodFailed)
		return
	}
}


func pushOperation(w http.ResponseWriter,r *http.Request)error{
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), allService)
	switch  {
	case err != nil :
		config.Log.Info("[%v] push json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	case allService.GetAllSrv():
		config.Log.Info("[%v] the service name have null", time.Now(),allService.SrvName)
		errResult.SendErrorResponse(w, config.ErrorSrvName)
		errs := fmt.Sprintf("the service name : %s have null", allService.SrvName)
		return errors.New(errs)
	default:
		err,s := allService.GetOneSrv()
		if err != nil {
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err
		}
		manager ,err := json.Marshal(s[0])
		normalResult.Resp = string(manager)
		normalResult.Code = 200
		result.NormalResponse(w, normalResult)
		return nil
	}
}