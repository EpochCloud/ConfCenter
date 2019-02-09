package main

import (
	"ConfCenter/router"
	"ConfCenter/config"
	"ConfCenter/initialization"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	variable := config.Variable{}
	variable.Initialize("debug","D:/project/src/ConfCenter/logcatlog")
	initialization.Initialize()

	router.Run()
}
