package engine

import (
	"ConfCenter/basic/util"
	"ConfCenter/config"
	"ConfCenter/basic"
	"encoding/json"
	"net/http"
	"errors"
	"time"
	"fmt"
	"ConfCenter/model/mysql"
)

/*
	get：查询
	post：插入
	patch：修改
*/

func OperationService(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		err, b := insertService(w, r)
		if err != nil || !b {
			return
		}
		return
	case r.Method == http.MethodGet:
		err := GetService(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodPatch:
		err, b := PatchService(w, r)
		if err != nil || !b {
			return
		}
		return
	case r.Method == http.MethodDelete:
		err := deleteSrv(w,r)
		if err != nil {
			return
		}
		return
	default:
		errResult.SendErrorResponse(w, config.ErrorMethodFailed)
		return
	}

}

//这里注意，服务名字是唯一的
func insertService(w http.ResponseWriter, r *http.Request) (error, bool) {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), service)
	switch {
	case err != nil:
		config.Log.Info("[%v] post json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err, false
		//这里要查询servicename在数据库中有没有
	case !service.GetService():
		config.Log.Info("[%v] the service name is  existed",time.Now(), service.ServiceName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return nil, false
	default:
		err := service.InsertService()
		if err != nil {
			config.Log.Error("[%v] insert opration err ",time.Now(), err)
			errResult.SendErrorResponse(w, config.DbError)
			return err, false
		}
		go do(w,service,r.Method)
		result.Response(w)
		return nil, true
	}
}

func GetService(w http.ResponseWriter, r *http.Request) error {
	err, s := service.GetAllService()
	if err != nil {
		config.Log.Error("[%v] insert opration err ",time.Now(), err)
		errResult.SendErrorResponse(w, config.DbError)
		return err
	}
	res := make(map[string]interface{}, 1)
	res["result"] = s

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

func PatchService(w http.ResponseWriter, r *http.Request) (error, bool) {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), service)
	switch {
	case err != nil:
		config.Log.Info("[%v] post json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err, false
	case service.GetService():
		config.Log.Info("[%v] the service name is  existed",time.Now(), service.ServiceName)
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return nil, false
	default:
		err := service.UpdateService()
		if err != nil {
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err, false
		}
		go do(w,service,r.Method)
		normalResult.Resp = "更新成功"
		normalResult.Code = 200
		result.NormalResponse(w, normalResult)
		return nil, true
	}
}


func  deleteSrv(w http.ResponseWriter,r *http.Request)error{
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()
	err := json.Unmarshal(body.Bytes(), service)
	switch  {
	case err != nil :
		config.Log.Info("[%v] delete json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	case service.GetService():
		config.Log.Info("[%v] the service name have null",time.Now(), service.ServiceName)
		errResult.SendErrorResponse(w, config.ErrorSrvName)
		errs := fmt.Sprintf("the service name : %s have null", service.ServiceName)
		return errors.New(errs)
	default:
		err := service.DeleteService()
		if err != nil {
			config.Log.Error("[%v] delete db err",time.Now(),err)
			errResult.SendErrorResponse(w, config.OperationDbErr)
			return err
		}
		go do(w,service,r.Method)
		result.Response(w)
		return nil
	}
}


func do(w http.ResponseWriter,service *mysql.Operation,method string){
	err, manager := getwayManager.GetConfiguration()
	defer func(){
		if err := recover();err != nil {
			config.Log.Error("operation ConfCenter to apigateway goroutine panic",err)
		}
	}()
	if err != nil {
		config.Log.Error("[%v] operation ConfCenter db to apigateway  err",time.Now(), err)
		errResult.SendErrorResponse(w, config.DbError)
		return
	}
	config.Log.Debug("[%v] the gatewaymanager is  %v",time.Now(),manager[0])
	config.Log.Debug("[%v] the service is %v",time.Now(),service)
	domain := fmt.Sprintf("%s%s:%s%s", scheme, manager[0].IntranetIp, manager[0].IntranetPort, manager[0].Serviceroute)
	go util.App(method, domain, service)
}