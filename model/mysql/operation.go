package mysql

import "ConfCenter/config"

type AllService struct {
	Id      uint64 `json:"id"`      //id
	Route   string `json:"route"`   //路由
	Ip      string `json:"ip"`      //ip
	Port    string `json:"port"`    //端口
	Srv     string `json:"srv"`     //所有服务的配置
	SrvName string `json:"srvname"` //服务名字
}

func NewAllService() *AllService {
	return &AllService{}
}

func (s *AllService) GetSrv() (err error, srv []*AllService) {
	err = config.Db.Select(&srv, "select * from allservice")
	if err != nil {
		return
	}
	return
}

func (s *AllService) InsertSrv() error {
	_, err := config.Db.Exec("insert into allservice(route,ip,port,srv,srvname) values(?,?,?,?,?)", s.Route, s.Ip, s.Port, s.Srv, s.SrvName)
	if err != nil {
		return err
	}
	return nil
}

func (s *AllService) GetAllSrv() bool {
	srv := make([]AllService, 0)
	config.Log.Debug("this is srv name", s.SrvName)
	err := config.Db.Select(&srv, "select * from allservice where srvname=?", s.SrvName)
	if err != nil {
		config.Log.Error("get service name err", err)
		return false
	}
	config.Log.Debug("this is srv", srv)
	if len(srv) == 0 {
		return true
	}
	return false
}

func (s *AllService) PatchSrv() error {
	_, err := config.Db.Exec("update allservice set route=?,ip=?,port=?,srv=? where srvname=?", s.Route, s.Ip, s.Port, s.Srv, s.SrvName)
	if err != nil {
		return err
	}
	return nil
}
