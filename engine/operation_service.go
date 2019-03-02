package engine

import (
	"ConfCenter/basic"
	"ConfCenter/basic/util"
	"ConfCenter/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Operation(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		err := getSrv(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodPost:
		err := insertSrv(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodPatch:
		err := patch(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodDelete:
		err := delService(w,r)
		if err != nil {
			return
		}
		return
	default:
		errResult.SendErrorResponse(w, config.ErrorMethodFailed)
		return
	}
}

func insertSrv(w http.ResponseWriter, r *http.Request) error {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), allService)
	switch {
	case err != nil:
		config.Log.Info("[%v] post json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	case !allService.GetAllSrv():
		config.Log.Info("[%v] the service name is  existed",time.Now(), allService.SrvName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		errs := fmt.Sprintf("the service name is : %s existed ", allService.SrvName)
		return errors.New(errs)
	default:
		err := allService.InsertSrv()
		if err != nil {
			config.Log.Error("[%v] insert opration err ", time.Now(),err)
			errResult.SendErrorResponse(w, config.DbError)
			return err
		}
		result.Response(w)
		return nil
	}
}

func getSrv(w http.ResponseWriter, r *http.Request) error {
	err, srv := allService.GetSrv()
	if err != nil {
		config.Log.Error("[%v] insert opration err ",time.Now(), err)
		errResult.SendErrorResponse(w, config.DbError)
		return err
	}
	res := make(map[string]interface{}, 1)
	res["result"] = srv

	massage, err := json.Marshal(res)
	if err != nil {
		config.Log.Info("[%v] get json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	}
	normalResult.Resp = string(massage)
	normalResult.Code = 200
	result.NormalResponse(w, normalResult)
	return nil
}

func patch(w http.ResponseWriter, r *http.Request) error {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), allService)
	switch {
	case err != nil:
		config.Log.Info("[%v] patch json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	case allService.GetAllSrv():
		config.Log.Info("[%v] the service name is  existed",time.Now(), allService.SrvName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		errs := fmt.Sprintf("the service name is : %s existed ", allService.SrvName)
		return errors.New(errs)
	default:
		err := allService.PatchSrv()
		if err != nil {
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err
		}
		normalResult.Resp = "更新成功"
		normalResult.Code = 200
		result.NormalResponse(w, normalResult)
		domain := fmt.Sprintf("%s%s:%s%s", scheme, allService.Ip, allService.Port, allService.Route)
		go util.App("POST", domain, allService)
		return nil
	}
}

//根据服务名字删除
func delService(w http.ResponseWriter, r *http.Request)error{
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), allService)
	switch  {
	case err != nil :
		config.Log.Info("[%v] delete json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	case allService.GetAllSrv():
		config.Log.Info("[%v] the service name have null",time.Now(), allService.SrvName)
		errResult.SendErrorResponse(w, config.ErrorSrvName)
		errs := fmt.Sprintf("the service name : %s have null", allService.SrvName)
		return errors.New(errs)
	default:
		err := allService.DeleteSrv()
		if err != nil {
			config.Log.Error("[%v] delete db err",time.Now(),err)
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err
		}
		result.Response(w)
		return nil
	}
}
