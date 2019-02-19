package main

import (
	"ConfCenter/router"
	"flag"
	"ConfCenter/initialization"
)

func main(){
	conf := flag.String("f", "./config/config.toml", "config file")
	flag.Parse()
	initialization.Initialize(*conf)
	router.Run()
}
