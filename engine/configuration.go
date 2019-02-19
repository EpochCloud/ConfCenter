package engine

import (
	"ConfCenter/basic"
	"ConfCenter/config"
	"encoding/json"
	"net/http"
	"ConfCenter/model/mysql"
)

func Configuration(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		err, b := insert(w, r)
		if err != nil || !b {
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

func insert(w http.ResponseWriter, r *http.Request) (error, bool) {
	body := basic.GetBody(w, r)
	defer func() {
		basic.Clean(w, r, body)
	}()

	err := json.Unmarshal(body.Bytes(), getwayManager)
	switch {
	case err != nil:
		config.Log.Info("post json Unmarshal err", err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err, false
		//不能为1，为1是代表以前填写过
	case getwayManager.Modification == 1:
		config.Log.Error("request Modification == 1")
		errResult.SendErrorResponse(w, config.PermissionError)
		return err, false
		// //进行验证验证以前最后一个为0的数据和现在填写的数据是否相同，相同就不行，如果这个验证成功就把上次的最后一个数据的modification设置为0
	case !validate(w, getwayManager):
		config.Log.Info("the configuration repetition")
		errResult.SendErrorResponse(w, config.ErrorRepeat)
		return err, false
	default:
		config.Log.Debug("result", getwayManager)
		err = getwayManager.SaveConfiguration()
		if err != nil {
			errResult.SendErrorResponse(w, config.DbError)
			config.Log.Error("save db", err)
			return err, false
		}
		result.Response(w)
		return nil, true
	}
}

func validate(w http.ResponseWriter, manager *mysql.GatewayManager) bool {
	err, m := manager.GetConfiguration()
	switch {
	case err != nil:
		config.Log.Error("get db err", err)
		return false
	case m[0].Ip == manager.Ip && m[0].Port == manager.Port && m[0].TimeOut == manager.TimeOut && m[0].LogPath == manager.LogPath && m[0].LogLevel == manager.LogLevel && m[0].Modification == manager.Modification && m[0].BufPool == manager.BufPool:
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
		config.Log.Error("get db err", err)
		errResult.SendErrorResponse(w, config.DbError)
		return err
	}
	m, err := json.Marshal(manager[0])
	if err != nil {
		config.Log.Info("get json Unmarshal err", err)
		errResult.SendErrorResponse(w, config.ErrorJsonFailed)
		return err
	}
	normalResult.Resp = string(m)
	normalResult.Code = 200
	result.NormalResponse(w, normalResult)
	return nil
}
