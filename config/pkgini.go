package config

import (
	"ConfCenter/log"
	"fmt"
)

//to be confirmed
func (variable *Variable)NewVeriable(level,logPath string)*Variable{
	variable.LogVariable.Init()
	variable.LogVariable.Set(level,logPath)
	return variable
}

func (variable *Variable)Initialize(level,logPath string){
	initLog(level,logPath)
}

func initLog(level,logpath string){
	var err error
	Log,err = log.New(level,logpath,0)
	if err != nil {
		fmt.Println("new log err",err)
		panic(err)
	}
}