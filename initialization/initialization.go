package initialization

import (
	"github.com/jmoiron/sqlx"
	"ConfCenter/log"
	"sync"
	"bytes"
	"ConfCenter/config"
)

type Initialization struct {
	Db *sqlx.DB
	BufPool *sync.Pool
}

func Initialize(){
	config.BufPool = NewInitialization().InitDb().BufPoolBasic().BufPool
}

func NewInitialization()*Initialization{
	return &Initialization{}
}

func (initialization *Initialization)InitDb()*Initialization{
	db ,err := sqlx.Open("mysql","root:123456@tcp(localhost:3306)/confcenter")
	if err != nil {
		log.Error("open mysql err",err)
		panic(err)
	}
	err = db.Ping()
	db.Close()
	if err != nil {
		log.Error("ping mysql err",err)
		panic(err)
	}
	initialization.Db = db
	return initialization
}

func (initialization *Initialization)BufPoolBasic()*Initialization{
	bufPool := &sync.Pool{
		New:MakeBuf,
	}
	for i := 0;i < 10; i ++ {
		bufPool.Put(bufPool.New())
	}
	initialization.BufPool = bufPool
	return initialization
}

func MakeBuf()interface{}{
	return bytes.NewBuffer(make([]byte,0,2048))
}