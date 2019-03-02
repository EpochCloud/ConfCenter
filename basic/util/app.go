package util

import (
	"net/http"
	"ConfCenter/config"
	"bytes"
	"encoding/json"
	"time"
)

func Proto(method,domain string,body *bytes.Buffer)(error,*http.Request){
	req,err := http.NewRequest(method,domain,body)
	if err != nil {
		config.Log.Error("[%v] http.request host err",time.Now(),err)
		return err,nil
	}
	req.Header.Add("Content-Type","application/json;charset=UTF-8")
	return nil,req
}

func App(method,domain string,b interface{}){
	body,err := json.Marshal(b)
	if err != nil {
		config.Log.Debug("[%v] app marshal body err",time.Now(),err)
		return
	}
	config.Log.Debug("[%v] method,host",time.Now(),method,domain,"body---",string(body))
	reqBody := bytes.NewBuffer(body)
	err,req := Proto(method,domain,reqBody)
	if err != nil {
		return
	}
	c := config.Client.Get()
	resp,err := c.(*http.Client).Do(req)
	if err != nil {
		config.Log.Error("[%v] resp err",time.Now(),err)
		return
	}
	defer func(c interface{}){
		resp.Body.Close()
		config.Client.Put(c)
		if err := recover();err != nil {
			config.Log.Error("[%v] this App goroutine err",time.Now(),err)
			return
		}
	}(c)
	if resp.StatusCode == 200 {
		config.Log.Debug("[%v] successfull about app method %s,host %s",time.Now(),method,domain)
		return
	}
	config.Log.Error("[%v] the resp code  is ",time.Now(),resp.StatusCode)
	return
}