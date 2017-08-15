package model

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	engine int
}

var engine *xorm.Engine

func getEngine()(*xorm.Engine){
	var err error
	if engine == nil{
		//engine, err = xorm.NewEngine("mysql", "root@/test_gin")
		engine, err = xorm.NewEngine("mysql", "etrip1:etrip!2017@tcp(123.56.120.16:3306)/test")
		if err != nil{
			panic(err)
		}
	}
	return engine
}
