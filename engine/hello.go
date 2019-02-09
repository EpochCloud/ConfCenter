package engine

import (
	"net/http"
	"io"
	"ConfCenter/basic"
	"encoding/json"
)

type HelloWorld struct {
	Result string  	`json:"result"`
}

func Hello(w http.ResponseWriter,r *http.Request){
	b := basic.GetBody(w,r)
	h := HelloWorld{
		Result:"hello world",
	}
	defer basic.Clean(w,r,b)
	helloWorld,_ := json.Marshal(h)
	io.WriteString(w,b.String()+string(helloWorld))
	return
}
