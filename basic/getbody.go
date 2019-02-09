package basic

import (
	"bytes"
	"io"
	"net/http"
	"ConfCenter/config"
)

func GetBody(w http.ResponseWriter,req *http.Request)*bytes.Buffer{
	r := config.BufPool.Get().(*bytes.Buffer)
	io.Copy(r,req.Body)
	return r
}

func Clean(w http.ResponseWriter,req *http.Request,r *bytes.Buffer){
	req.Body.Close()
	r.Reset()
	config.BufPool.Put(r)
}


