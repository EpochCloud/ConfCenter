package engine

import (
	"ConfCenter/basic"
	"ConfCenter/config"
	"encoding/json"
	"net/http"
	"ConfCenter/model/mysql"
	"fmt"
	"ConfCenter/basic/util"
	"time"
	"errors"
)

func Configuration(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		err := insert(w, r)
		if err != nil {
			return
		}
		return
	case r.Method == http.MethodGet:
		err := get(w, r)
		if err != nil {
			return
		}
		return
	default:
		errResult.SendErrorResponse(w, config.ErrorMethodFailed)
		return
	}
}

func insert(w http.ResponseWriter, r *http.Request) error {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()

	err := json.Unmarshal(body.Bytes(), getwayManager)
	switch {
	case err != nil:
		config.Log.Info("[%v] post json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
		//不能为1，为1是代表以前填写过
	case getwayManager.Modification == 1:
		config.Log.Error("[%v] request Modification == 1",time.Now())
		errResult.SendErrorResponse(w, config.PermissionError)
		return err
		// //进行验证验证以前最后一个为0的数据和现在填写的数据是否相同，相同就不行，如果这个验证成功就把上次的最后一个数据的modification设置为0
	case !validate(w, getwayManager):
		config.Log.Info("[%v] the configuration repetition",time.Now())
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return errors.New("the configuration repetition")
	default:
		config.Log.Debug("[%v] result",time.Now(), getwayManager)
		err = getwayManager.SaveConfiguration()
		if err != nil {
			errResult.SendErrorResponse(w, config.DbError)
			config.Log.Error("[%v] save db",time.Now(), err)
			return err
		}
		domain := fmt.Sprintf("%s%s:%s%s", scheme, getwayManager.IntranetIp, getwayManager.IntranetPort, getwayManager.Managerroute)
		go util.App("POST", domain, getwayManager)
		result.Response(w)
		return nil
	}
}

func validate(w http.ResponseWriter, manager *mysql.GatewayManager) bool {
	err, m := manager.GetConfiguration()
	switch {
	case err != nil:
		config.Log.Error("[%v] get db err",time.Now(), err)
		return false
	case *m[0] == *manager:
		return false
	default:
		err := manager.AlterConfiguration(m[0].Id)
		if err != nil {
			return false
		}
		return true
	}
}

//这个是直接展示
func get(w http.ResponseWriter, r *http.Request) error {
	err, manager := getwayManager.GetConfiguration()
	if err != nil {
		config.Log.Error("[%v] get db err",time.Now(), err)
		errResult.SendErrorResponse(w, config.DbError)
		return err
	}
	m, err := json.Marshal(manager[0])
	if err != nil {
		config.Log.Info("[%v] get json Unmarshal err",time.Now(), err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	}
	normalResult.Resp = string(m)
	normalResult.Code = 200
	result.NormalResponse(w, normalResult)
	return nil
}