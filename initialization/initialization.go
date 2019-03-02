package initialization

import (
	"ConfCenter/config"
	"ConfCenter/log"
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Initialization struct {
	db             *sqlx.DB
	bufPool, clent *sync.Pool
	Config         config.Config
}

func Initialize(conf string) {
	newInitialization := newInitialization().initConfig(conf).initlog().initDb().bufPoolBasic().serverClient()
	config.BufPool = newInitialization.bufPool
	config.Db = newInitialization.db
	config.Conf = &newInitialization.Config
	config.Client = newInitialization.clent
}

func newInitialization() *Initialization {
	return &Initialization{}
}

func (initialization *Initialization) initDb() *Initialization {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%v:%v)/%v", initialization.Config.MysqlConf.Name, initialization.Config.MysqlConf.PassWord, initialization.Config.MysqlConf.Ip, initialization.Config.MysqlConf.Port, initialization.Config.MysqlConf.Database)
	config.Log.Debug("[%v] mysql config is :%v",time.Now(), dataSourceName)
	db, err := sqlx.Open(initialization.Config.MysqlConf.DriverName, dataSourceName)
	if err != nil {
		log.Error("open mysql err", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Error("ping mysql err", err)
		panic(err)
	}
	initialization.db = db
	return initialization
}

func (initialization *Initialization) bufPoolBasic() *Initialization {
	bufPool := &sync.Pool{
		New: MakeBuf,
	}
	for i := 0; i < 10; i++ {
		bufPool.Put(bufPool.New())
	}
	initialization.bufPool = bufPool
	return initialization
}

func MakeBuf() interface{} {
	return bytes.NewBuffer(make([]byte, 0, 2048))
}

func (initialization *Initialization) initlog() *Initialization {
	initLog(initialization.Config.LogConf.LogLevel, initialization.Config.LogConf.LogPath)
	return initialization
}

func initLog(level, logpath string) {
	var err error
	config.Log, err = log.New(level, logpath, 0)
	if err != nil {
		log.Error("new log err", err)
		panic(err)
	}
}

func (initialization *Initialization) initConfig(conf string) *Initialization {
	configBytes, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Error("ioutil readfile config err:%v", err)
		panic(err)
	}
	if _, err := toml.Decode(string(configBytes), &initialization.Config); err != nil {
		log.Error("toml decode err ", err)
		panic(err)
	}
	log.Debug("all config is %v", initialization.Config)
	return initialization
}

func (initialization *Initialization) serverClient() *Initialization {
	client := &sync.Pool{
		New: MakeClient,
	}
	for i := 0; i < 10; i++ {
		client.Put(client.New())
	}
	initialization.clent = client
	return initialization
}

func MakeClient() interface{} {
	return &http.Client{}
}
