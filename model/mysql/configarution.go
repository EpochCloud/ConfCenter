package mysql

import (
	"ConfCenter/config"
)

type GatewayManager struct {
	Id           uint64 `json:"id"`
	Ip           string `json:"ip"`           //api服务的ip
	Port         string `json:"port"`         //api服务的端口
	TimeOut      int    `json:"timeout"`      //api设置的超时时间
	LogLevel     string `json:"loglevel"`     //日志的级别
	LogPath      string `json:"logpath"`      //日志的路径
	Modification uint64 `json:"modification"` //是否被覆盖，覆盖了是1，不覆盖是0
	BufPool      int    `json:"bufpool"`      //buf池子的容量
}

func NewGatewayManager() *GatewayManager {
	return &GatewayManager{}
}

func (gatewayManager *GatewayManager) SaveConfiguration() error {
	_, err := config.Db.Exec("insert into configuration(ip,port,timeout,loglevel,logpath,modification,bufpool) values(?,?,?,?,?,?,?)", gatewayManager.Ip, gatewayManager.Port, gatewayManager.TimeOut, gatewayManager.LogLevel, gatewayManager.LogPath, gatewayManager.Modification, gatewayManager.BufPool)
	if err != nil {
		config.Log.Error("save configuration err", err)
		return err
	}
	return nil
}

func (gatewayManager *GatewayManager) GetConfiguration() (err error, Manager []*GatewayManager) {
	err = config.Db.Select(&Manager, "select  * from configuration where modification=0")
	if err != nil {
		config.Log.Error("get configuration err", err)
		return err, nil
	}
	return nil, Manager
}

func (gatewayManager *GatewayManager) AlterConfiguration(id uint64) error {
	_, err := config.Db.Exec("update configuration set modification=? where id=?", 1, id)
	if err != nil {
		config.Log.Error("update configuration err", err)
		return err
	}
	return nil
}
