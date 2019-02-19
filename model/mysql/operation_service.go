package mysql

import (
	"ConfCenter/config"
	"encoding/json"
)

/*
	post:
		插入数据，这里插入数据的时候要做一个对比，即如果服务名字一样，那么就说明操作失误
	patch:
		修改操作，这里修改是根据服务名字这个唯一的不重复的服务名字来进行修改的
		这里把修改的数据单独保存一张表，这个表要多有一个id选项，等于上面标的id
	get:
		这里是查看所有的服务注册的相信信息，这里要用limit分页

*/
type Operation struct {
	Id          uint64   `json:"id"`    //主id
	Route       string   `json:"route"` //路由
	Service     *Service `json:"service"`
	ServiceName string   `json:"servicename"` //服务名字
}

type Operations struct {
	Id          uint64 `json:"id"`    //主id
	Route       string `json:"route"` //路由
	Service     string `json:"service"`
	ServiceName string `json:"servicename"` //服务名字
}

type Service struct {
	ServiceAddr []string `json:"serviceaddr"` //服务地址  [ip:port]
	//RegisterName string   `json:"registername"` //谁注册的服务  这里的名字是登录的名字，不能让人填写，这里先空着，等登录注册完成之后再说补充
	RegisterTime string `json:"registertime"` //注册时间
	//AltTime      string   `json:"alttime"`      //修改时间  这里的名字是登录的名字，不能让人填写，这里先空着，等登录注册完成之后再说补充
	AltReason string `json:"altreason"` //修改原因
}

func NewService() *Operation {
	return &Operation{
		Service: &Service{
			ServiceAddr: make([]string, 0, 10),
		},
	}
}

func (service *Operation) InsertService() error {
	s, err := json.Marshal(service.Service)
	if err != nil {
		config.Log.Error("insert db err", err)
		return err
	}
	config.Log.Debug("insert service.service", string(s))
	_, err = config.Db.Exec("insert into service(route,service,servicename) values(?,?,?)", service.Route, string(s), service.ServiceName)
	if err != nil {
		return err
	}
	return nil
}

func (service *Operation) GetService() bool {
	s := make([]*Operations, 0)
	err := config.Db.Select(&s, "select * from service where servicename=?", service.ServiceName)
	if err != nil {
		config.Log.Error("get service name err", err)
		return false
	}
	if len(s) == 0 {
		return true
	}
	return false
}

func (service *Operation) GetAllService() (error, []*Operations) {
	s := make([]*Operations, 0, 10)
	err := config.Db.Select(&s, "select * from service")
	if err != nil {
		config.Log.Error("get service name err", err)
		return err, nil
	}
	return nil, s
}

func (service *Operation) UpdateService() error {
	s, err := json.Marshal(service.Service)
	if err != nil {
		config.Log.Error("insert db err", err)
		return err
	}
	_, err = config.Db.Exec("update service set router=?,service=? where servicename=?", service.Route, string(s), service.ServiceName)
	if err != nil {
		config.Log.Error("update service err", err)
		return err
	}
	return nil
}
