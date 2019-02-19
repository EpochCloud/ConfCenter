package util

import (
	"net/http"
	"ConfCenter/config"
	"bytes"
	"encoding/json"
)

func Proto(method,domain string,body *bytes.Buffer)(error,*http.Request){
	req,err := http.NewRequest(method,domain,body)
	if err != nil {
		config.Log.Error("http.request host err",err)
		return err,nil
	}
	req.Header.Add("Content-Type","application/json;charset=UTF-8")
	return nil,req
}

func App(method,domain string,b interface{}){
	body,err := json.Marshal(b)
	if err != nil {
		config.Log.Debug("app marshal body err",err)
		return
	}
	config.Log.Debug("method,host",method,domain,"body---",body)
	reqBody := bytes.NewBuffer(body)
	err,req := Proto(method,domain,reqBody)
	if err != nil {
		return
	}
	c := config.Client.Get()
	resp,err := c.(*http.Client).Do(req)
	if err != nil {
		config.Log.Error("resp err",err)
		return
	}
	defer func(){
		resp.Body.Close()
		config.Client.Put(c)
		if err := recover();err != nil {
			config.Log.Error("this App goroutine err",err)
			return
		}
	}()
	if resp.StatusCode == 200 {
		config.Log.Debug("successfull about app method %s,host %s",method,domain)
		return
	}
	config.Log.Error("the resp code  is ",resp.StatusCode)
	return
}