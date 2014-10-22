package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	// db.Init()
	db_host, _ := r.Config.String("db.host")
	db_port, _ := r.Config.String("db.port")
	db_user, _ := r.Config.String("db.user")
	db_pass, _ := r.Config.String("db.password")
	db_name, _ := r.Config.String("db.name")
	db_protocol, _ := r.Config.String("db.protocol")

	connectionString := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s", db_user, db_pass, db_protocol, db_host, db_port, db_name)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("1111111")
	}

	// construct a gorp DbMap
	Dbm := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TraceOn("[gorp]", r.INFO)

	count, _ := Dbm.SelectInt("select count(*) from Hotel")
	log.Println("Row count - should be zero:", count)
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
